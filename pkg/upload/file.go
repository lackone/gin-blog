package upload

import (
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/util"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

func GetFileName(name string) string {
	ext := GetFileExt(name)
	filename := strings.TrimSuffix(name, ext)
	return util.EncodeMd5(filename) + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func CheckSavePath(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

func CheckExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckSize(t FileType, f multipart.File) bool {
	all, _ := io.ReadAll(f)
	size := len(all)
	switch t {
	case TypeImage:
		if size <= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, mode os.FileMode) error {
	err := os.MkdirAll(dst, mode)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(f *multipart.FileHeader, dst string) error {
	open, err := f.Open()
	if err != nil {
		return err
	}
	defer open.Close()
	create, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(create, open)
	return err
}
