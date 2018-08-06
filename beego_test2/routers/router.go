package routers

import (
	"github.com/astaxie/beego"
	"github.com/deepaksinghkushwah/beego_test2/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/mydata", &controllers.MainController{}, "get:Mydata")
}
