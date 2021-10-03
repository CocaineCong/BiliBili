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

type UserUpdate struct {
	Name string `form:"name" json:"name" bind:"required"`
	Gender int `form:"gender" json:"gender" `
	Birthday string `form:"birthday" json:"birthday" time_format:"2006-01-02"`
	Sign string `form:"sign" json:"sign" `
}

type UserInfo struct {

}

type UserSearch struct {
	UserName string `form:"user_name" json:"user_name"`
}

func (service *UserRegister) Register() serializer.Response {
	var user model.User
	var count int
	code := e.SUCCESS
	if !utils.VerifyEmailFormat(service.Email) {
		code = e.InvalidParams
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"不是正确的邮箱格式",
		}
	}
	model.DB.Model(&model.User{}).Where("email = ?",service.Email).Count(&count)
	if count>0 {
		code = e.InvalidParams
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

func (service *UserLogin) Login() serializer.Response {
	code := e.SUCCESS
	var user model.User
	if !utils.VerifyEmailFormat(service.Email) {
		code = e.InvalidParams
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"不是正确的邮箱格式",
		}
	}
	model.DB.Model(&model.User{}).Where("email = ?",service.Email).First(&user)
	if user.ID == 0{
		code = e.InvalidParams
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"用户不存在",
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(service.Password)) ; err != nil {
		code = e.InvalidParams
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

func (service *UserUpdate) Update(id uint) serializer.Response {
	code := e.SUCCESS
	err := model.DB.Model(model.User{}).Where("id = ?",id).
		Updates(map[string]interface{}{
			"user_name":service.Name,"gender":service.Gender,
			"birthday":service.Birthday,"sign":service.Sign}).Error
	if err!=nil {
		code = e.ERROR
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
			Data:"修改信息失败",
		}
	}
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:"修改信息成功",
	}
}

func (service *UserInfo) Show(id uint) serializer.Response {
	code := e.SUCCESS
	var user model.User
	model.DB.Model(&model.User{}).First(&user,id)
	return serializer.Response {
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildUser(user),
	}
}

func (service *UserSearch) Search() serializer.Response {
	code := e.SUCCESS
	var user []model.User
	model.DB.Model(&model.User{}).
		Where("user_name LIKE ?","%"+service.UserName+"%").
		Find(&user)
	return serializer.Response{
		Status:code,
		Msg:e.GetMsg(code),
		Data:serializer.BuildUsers(user),
	}
}