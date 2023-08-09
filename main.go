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
	//设置静态资源路径 https://www.cnblogs.com/aaronthon/p/12802591.html
	r.Static("/res", "./res")
	//配置ssl证书
	r.RunTLS(":443", "", "")

	//自定义 http配置
	s := http.Server{
		Addr:              ":443",
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	//index 首页
	r.GET("/", ctr.Index)

	//启动 webserver http服务
	//err := s.ListenAndServe()

	//启动 https服务 并加载ssl证书 https://www.jianshu.com/p/01057d2c37e4
	err := s.ListenAndServeTLS("", "")
	if err != nil {
		// Gin 启动失败 终止程序
		panic(err)
	}

}
