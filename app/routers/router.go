package routers

import (
	"bytes"
	"github.com/admin/app/pkg/e"
	"github.com/admin/app/pkg/setting"
	"github.com/admin/app/routers/api/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)


func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//加载模板
	router.LoadHTMLFiles("app/templates/sys/index.html")
	//获取当前文件的相对路径
	router.StaticFile("/logo.png", "./app/templates/sys/logo.png")
	router.Static("static", "./app/static")
	//设置入口页
	router.GET("/", func(i *gin.Context) {
		i.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/user/login", func(i *gin.Context) {
		i.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/user/register", func(i *gin.Context) {
		i.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	//测试接口
	router.GET("/test", func(c *gin.Context) {
		json := make(map[string]interface{})
		json["time"] = time.Now().Format("2006-01-02 15:04:05")
		c.JSON(200, e.ResponseJson("操作成功",json,true))
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, e.NotFound())
	})

	//v1 路由
	apiV1 := router.Group("/api/v1")
	{
		//项目信息
		apiV1.GET("/project", func(c *gin.Context) {
			jsons := make(map[string]interface{})
			var bt bytes.Buffer
			if strings.Contains(c.Request.Proto,"https") {
				bt.WriteString("https://")
			} else {
				bt.WriteString("http://")
			}
			bt.WriteString(c.Request.Host)
			bt.WriteString("/static/images/")
			bt.WriteString(setting.LoadProject().Logo)
			jsons["name"] = setting.LoadProject().Name
			jsons["desc"] = setting.LoadProject().Desc
			jsons["logo"] = bt.String()
			jsons["developer"] = setting.LoadProject().Developer
			c.JSON(200, e.ResponseJson("操作成功",jsons,true))
		})
		//登录
		apiV1.POST("/user/login", user.Login)
	}

	return router
}
