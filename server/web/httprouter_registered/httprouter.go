/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:02 2019-09-16
 */
package httprouter_registered

import (
	router3 "easydevops/client/api/router"
	"github.com/julienschmidt/httprouter"
)

func RegisterHttprouter() *httprouter.Router {
	router := httprouter.New()

	router3.RegisterRouter(router)

	return router
}
