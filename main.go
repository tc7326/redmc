package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"redmc/ctr"
	"time"
)

func main() {

	//配置日志格式
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix("[LOG] ")

	//Gin的默认中间件 带日志和异常回调
	r := gin.Default()
	//设置html模板地址
	r.LoadHTMLGlob("html/*")
	//配置图标 https://blog.csdn.net/lt326030434/article/details/113058241
	r.StaticFile("/favicon.ico", "./icon.png")
	//设置静态资源路径 https://www.cnblogs.com/aaronthon/p/12802591.html
	r.Static("/res", "./res")

	//自定义 http配置
	s := http.Server{
		Addr:              ":443",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	//index 首页
	r.GET("/", ctr.Index)

	//启动 web https服务 并加载ssl证书
	err := s.ListenAndServeTLS("ssl/redcraft.cn.pem", "ssl/redcraft.cn.key")
	if err != nil {
		// Gin 启动失败 终止程序
		panic(err)
	}

}
