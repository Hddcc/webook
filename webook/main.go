package main

import (
	"net/http"
	"time"
	"webook/internal/repository"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)

	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("这是第一个middleware")
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	u := &web.UserHandler{}
	u.RegisterRoutes(server)
	server.Run(":8089")
}
