package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_sugared/routers/api/pic"
	"go_sugared/routers/api/room"
	"go_sugared/routers/api/user"
)

func InitRouter() *gin.Engine {
	fmt.Println("init routers")
	gin.SetMode("debug")
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	return addRoute(engine)
}

func addRoute(engine *gin.Engine) *gin.Engine {
	apiRoom := engine.Group("/api/room")
	{
		//获取全部房源信息
		apiRoom.GET("/get", room.GetRooms)
		//获取具体房源信息
		apiRoom.POST("/getDetail", room.GetRoomDetail)
		//新增房源信息
		apiRoom.POST("/add", room.AddRooms)
		//编辑房源信息
		apiRoom.POST("/update", room.UpdateRooms)
		//删除房源信息
		apiRoom.POST("/delete", room.DeleteRooms)
	}

	apiPic := engine.Group("api/pic")
	{
		//test
		apiPic.GET("/test", pic.Test)
		//新增房源pic
		apiPic.POST("/add", pic.AddPic)
		//删除房源图片信息
		apiPic.POST("/delete", pic.DeletePic)
	}

	apiUser := engine.Group("api/user")
	{
		//test
		apiUser.GET("/test", user.Test)
		//用户注册登录
		apiUser.POST("/login", user.UserLogin)
		//更新用户信息
		apiUser.POST("/update", user.UserUpdate)
	}

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return engine
}
