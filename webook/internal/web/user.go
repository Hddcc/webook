package web

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// UserHandler 定义所有与user有关的路由
type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")
	u.RegisterRoutesV1(api.Group("/v1"))
}

// RegisterRoutesV1 注册用户接口。
func (u *UserHandler) RegisterRoutesV1(v1 *gin.RouterGroup) {
	v1.POST("/users", u.SignUp)
	v1.POST("/sessions", u.Login)
	v1.GET("/users/me", u.Profile)
	v1.PATCH("/users/me", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok, err := regexp.Match(emailRegexPattern, []byte(req.Email))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "邮箱格式有误")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusBadRequest, "两次输入的密码不一致")
		return
	}
	ok, err = regexp.Match(passwordRegexPattern, []byte(req.Password))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "密码必须大于8位、包含数字和特殊字符")
		return
	}
	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v", req)
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
