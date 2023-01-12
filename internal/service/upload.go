package service

import (
	"errors"
	"github.com/lackone/gin-blog/global"
	"github.com/lackone/gin-blog/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name string
	Url  string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	name := upload.GetFileName(fileHeader.Filename)
	path := upload.GetSavePath()
	dst := path + "/" + name
	if !upload.CheckExt(fileType, name) {
		return nil, errors.New("类型不允许")
	}
	if upload.CheckSavePath(path) {
		err := upload.CreateSavePath(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if !upload.CheckSize(fileType, file) {
		return nil, errors.New("文件过大")
	}
	if upload.CheckPermission(path) {
		return nil, errors.New("禁止访问")
	}
	err := upload.SaveFile(fileHeader, dst)
	if err != nil {
		return nil, err
	}
	return &FileInfo{
		Name: name,
		Url:  global.AppSetting.UploadServerUrl + "/" + name,
	}, nil
}
