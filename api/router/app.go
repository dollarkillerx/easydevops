/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:06 2019-09-16
 */
package router

import (
	"easydevops/api/controller"
	"github.com/julienschmidt/httprouter"
)

func RegisterRouter(app *httprouter.Router) {
	app.POST("/update", controller.Rundevops)
}
