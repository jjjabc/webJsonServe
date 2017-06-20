package webJsonServe

import (
       "encoding/json"
       "github.com/astaxie/beego"
)

const (
	FAILED = "failed"
	SUCCESS = "success"
)

type RespJSON struct {
	Status string
	Des    string
	Data   *json.RawMessage
}

func ServeFailed(c *beego.Controller, status int, des string) {
	rJSON := RespJSON{}
	rJSON.Status = FAILED
	rJSON.Des = des
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = rJSON
	c.ServeJSON()
	return
}
func ServeSuccess(c *beego.Controller, des string, data interface{}) {
	rJSON := RespJSON{}
	rJSON.Status = SUCCESS
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

	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = rJSON
	c.ServeJSON()
}
