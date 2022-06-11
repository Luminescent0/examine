package api

import (
	"examine/model"
	"examine/service"
	"examine/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	err := service.UsernameIsExist(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tool.RespErrorWithDate(ctx, "用户不存在")
			return
		}
		tool.RespInternalError(ctx)
		return
	}
	flag, err := service.IsPasswordCorrect(username, password)
	if err != nil {
		fmt.Println("judge password correct err :", err)
		tool.RespInternalError(ctx)
		return
	}
	if !flag {
		tool.RespErrorWithDate(ctx, "密码错误")
		return
	}
	token, err1 := CreateToken(username)
	if err1 != nil {
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, token)
	return
}
func register(ctx *gin.Context) {
	username, password := verify(ctx)
	if username == "存在非法输入" {
		tool.RespErrorWithDate(ctx, "用户名格式有误(min=4,max=10)")
		return
	}
	if password == "" {
		tool.RespErrorWithDate(ctx, "密码格式有误(min=6,max=16)")
		return
	}
	user := model.User{
		Username: username,
		Password: password,
	}
	flag, err := service.IsRepeatUsername(username)
	if err != nil {
		fmt.Println("judge repeat username err:", err)
		tool.RespInternalError(ctx)
		return
	}
	if flag {
		tool.RespErrorWithDate(ctx, "用户名已经存在")
		return
	}

	err = service.Register(user)
	if err != nil {
		fmt.Println("register err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessfulWithDate(ctx, "注册成功")
}
func verify(ctx *gin.Context) (string, string) { //验证非法输入
	validate := validator.New() //创建验证器
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	u := model.User{Id: 0, Username: username, Password: password}

	err := validate.Struct(u)
	fmt.Println(err)
	if err != nil {
		return "存在非法输入", ""
	}
	return username, password

}
