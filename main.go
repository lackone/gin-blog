package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/internal/model"
	"github.com/lackone/gin-blog/internal/routers"
	"github.com/lackone/gin-blog/pkg/logger"
	"github.com/lackone/gin-blog/pkg/setting"
	"github.com/lackone/gin-blog/pkg/trace"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port string //端口
	mode string //模式
	conf string //配置文件路径
)

func init() {
	err := InitFlag()
	if err != nil {
		log.Fatalln(err)
	}
	err = InitSetting()
	if err != nil {
		log.Fatalln(err)
	}
	err = InitLogger()
	if err != nil {
		log.Fatalln(err)
	}
	err = InitDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = InitTrace()
	if err != nil {
		log.Fatalln(err)
	}
}

// @title 博客
// @version 1.0
// @description gin写的博客
// @termsOfService https://github.com/lackone/gin-blog
func main() {
	router := routers.NewRouter()
	s := http.Server{
		Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 32,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln("server listen err:", err)
		}
	}()

	//监控SIGINT和SIGTERM信号，实现优雅关机
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shut down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln("shut down server err:", err)
	}

	log.Println("server exit...")
}

// 初始化配置
func InitSetting() error {
	setting, err := setting.NewSetting(strings.Split(conf, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JwtSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if mode != "" {
		global.ServerSetting.RunMode = mode
	}
	return nil
}

// 初始化数据库
func InitDB() error {
	var err error
	global.DB, err = model.NewDb(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func InitLogger() error {
	logFile := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  logFile,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

// 初始化链路追踪
func InitTrace() error {
	var err error
	global.Trace, err = trace.NewTrace(global.ServerSetting.ServiceName, "http://localhost:14268/api/traces", global.ServerSetting.RunMode, 1)
	if err != nil {
		return err
	}
	return nil
}

// 初始化命令行参数
func InitFlag() error {
	flag.StringVar(&port, "port", "", "端口")
	flag.StringVar(&mode, "mode", "", "运行模式")
	flag.StringVar(&conf, "conf", "configs/", "配置文件路径")
	flag.Parse()
	return nil
}
