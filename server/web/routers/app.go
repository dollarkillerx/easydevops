/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:06 2019-09-16
 */
package routers

import (
	"easydevops/server/web/controller"
	"github.com/julienschmidt/httprouter"
)

func RegisterRouter(app *httprouter.Router) {
	app.POST("/upfile", controller.UpFile)  // cli端 上传文件  需要秘钥才能上传

	app.POST("/getfile",controller.GetFile)  //  需要秘钥才能获取
}
