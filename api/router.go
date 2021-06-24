package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tongren/api/v1"
	"tongren/middleware"
)

//主路由
func InitRouter() http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.Cors())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	//r.LoadHTMLGlob("dist/*.html")    // 添加入口index.html
	//r.LoadHTMLFiles("dist/*/*")	// 添加资源路径
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	r.Static("/css", "./dist/css") 	// 添加css
	r.Static("/js", "./dist/js") 	// 添加js
	r.Static("/img", "./dist/img") 	// 添加img
	r.StaticFile("/", "dist/index.html")  //前端接口

	app_v1 := r.Group("/api/v1")
	app_v1.Use(gin.Recovery())
	{
		app_v1.GET("/capt",v1.GetCapt)
		app_v1.POST("/capt",v1.CheckCapt)
		app_v1.POST("/register", v1.Register)
	}

	return r
}
