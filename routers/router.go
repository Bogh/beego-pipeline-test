package routers

import (
	"github.com/bogh/beego-pipeline-test/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
