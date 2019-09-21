/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:02 2019-09-16
 */
package httprouter_registered

import (
	"easydevops/server/web/routers"
	"github.com/julienschmidt/httprouter"
)

func RegisterHttprouter() *httprouter.Router {
	router := httprouter.New()

	routers.RegisterRouter(router)

	return router
}
