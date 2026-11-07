package middleware

import "github.com/gin-gonic/gin"

type LoginMiddlewareBuilder struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) { 
	// 不需要登录校验的
		if ctx.Request.URL.Path == "/api/v1/sessions" || ctx.Request.URL.Path == "/api/v1/users" {
			return
		}
		sess := sessions.Default(ctx)
		userId := sess.Get("userId")
		if userId == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}