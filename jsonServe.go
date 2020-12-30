package webJsonServe

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/jjjabc/webJsonServe/models"
)

func ServeFailed(c *beego.Controller, status int, des string) (rJSON models.RespJSON) {
	rJSON.Status = models.FAILED
	rJSON.Des = des
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = rJSON
	c.ServeJSON()
	return
}
func ServeSuccess(c *beego.Controller, des string, data interface{}) (rJSON models.RespJSON) {
	rJSON.Status = models.SUCCESS
	rJSON.Des = des
	if (data == nil) {
		rJSON.Data = nil
	} else {
		b, err := json.Marshal(data)
		if err != nil {
			ServeFailed(c, 500, err.Error())
			return
		}
		raw := json.RawMessage(b)
		rJSON.Data = &raw
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = rJSON
	c.ServeJSON()
	return
}
