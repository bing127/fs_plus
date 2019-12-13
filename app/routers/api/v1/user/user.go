package user

import (
	"encoding/json"
	"fmt"
	"github.com/admin/app/pkg/e"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

/*
登录
 */
func Login(c *gin.Context)  {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	err :=  json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(200,e.ResponseJson("操作成功",nil,false))
	}
	fmt.Println(user["username"])
	fmt.Println(user["password"])
	//c.JSON(200, e.ResponseJson("操作成功",nil,true))
}