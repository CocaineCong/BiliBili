package utils

import (
	"BiliBili.com/pkg/e"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"mime/multipart"
)

// 封装上传图片到七牛云然后返回状态和图片的url
func UploadToQiNiu(file multipart.File ,fileSize int64) (int,string) {
	var AccessKey =  viper.GetString("qiniu.AccessKey")
	var SerectKey = viper.GetString("qiniu.SerectKey")
	var Bucket = viper.GetString("qiniu.Bucket")
	var ImgUrl = viper.GetString("qiniu.QiniuServer")
	putPlicy := storage.PutPolicy{
		Scope:Bucket,
	}
	mac := qbox.NewMac(AccessKey,SerectKey)
	upToken := putPlicy.UploadToken(mac)
	cfg := storage.Config{
		Zone : &storage.ZoneHuanan,
		UseCdnDomains : false,
		UseHTTPS : false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil {
		code := e.ErrorUploadFile
		return code , err.Error()
	}
	url := ImgUrl + ret.Key
	return 200,url
}
