//author:xunj
//createDate:2019/12/19 下午12:35
//desc:$END$
package servegin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jjjabc/webJsonServe/models"
)

func ServeFailed(c *gin.Context, status int, des string) {
	rJSON := models.RespJSON{}
	rJSON.Status = models.FAILED
	rJSON.Des = des
	c.JSON(status,rJSON)
	return
}
func ServeSuccess(c *gin.Context, des string, data interface{}) {
	rJSON := models.RespJSON{}
	rJSON.Status = models.SUCCESS
	rJSON.Des = des
	if (data == nil) {
		rJSON.Data = nil
	} else {
		b, err := json.Marshal(data)
		if err != nil {
			ServeFailed(c,500,err.Error())
			return
		}
		raw := json.RawMessage(b)
		rJSON.Data = &raw
	}
	c.JSON(200,rJSON)
}
