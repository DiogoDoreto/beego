// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/DiogoDoreto/beego for the canonical source repository
// @license     http://github.com/DiogoDoreto/beego/blob/master/LICENSE
// @authors     astaxie

package main

import (
	"github.com/DiogoDoreto/beego"
	"github.com/DiogoDoreto/beego/example/beeapi/controllers"
)

//		Objects

//	URL					HTTP Verb				Functionality
//	/object				POST					Creating Objects
//	/object/<objectId>	GET						Retrieving Objects
//	/object/<objectId>	PUT						Updating Objects
//	/object				GET						Queries
//	/object/<objectId>	DELETE					Deleting Objects

func main() {
	beego.RESTRouter("/object", &controllers.ObjectController{})
	beego.Run()
}
