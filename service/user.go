package service

import (
	"BiliBili.com/model"
	"BiliBili.com/pkg/e"
	"BiliBili.com/pkg/utils"
	"BiliBili.com/serializer"
	"golang.org/x/crypto/bcrypt"
)

type UserRegister struct {
	UserName string `json:"user_name" form:"user_name" bind:"required"`
	Email string `json:"email" form:"email" bind:"required"`
	Password string `json:"password" form:"password" bind:"required"`
	//Code string `json:"code" form:"code" bind:"required"`
	Code string `json:"code" form:"code"`
}

type UserLogin struct {
	Email string `json:"email" form:"email" bind:"required"`
	Password string `json:"password" form:"password" bind:"required"`
}

func (service *UserRegister) Register() serializer.Response {
	var user model.User
	var count int
	code := e.SUCCESS
	if !utils.VerifyEmailFormat(service.Email) {
		code = e.INVALID_PARAMS
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"不是正确的邮箱格式",
		}
	}
	model.DB.Model(&model.User{}).Where("email = ?",service.Email).Count(&count)
	if count>0 {
		code = e.INVALID_PARAMS
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"这个邮箱已经存在了",
		}
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(service.Password),bcrypt.DefaultCost)
	user = model.User{
		Email:service.Email,
		UserName:service.UserName,
		Password : string(hashedPassword),
	}
	model.DB.Create(&user)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"创建用户成功",
	}
}

func (service *UserLogin) UserLogin() serializer.Response {
	code := e.SUCCESS
	var user model.User
	if !utils.VerifyEmailFormat(service.Email) {
		code = e.INVALID_PARAMS
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"不是正确的邮箱格式",
		}
	}
	model.DB.Model(&model.User{}).Where("email = ?",service.Email).First(&user)
	if user.ID == 0{
		code = e.INVALID_PARAMS
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"用户不存在",
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(service.Password)) ; err != nil {
		code = e.INVALID_PARAMS
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"密码不正确",
		}
	}
	user.Authority = 0
	token ,_ := utils.ReleaseToken(user)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.TokenData{User:serializer.BuildUser(user),Token:token},
	}
}
