/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:01 2019-09-16
 */
package main

import (
	httprouter_registered2 "easydevops/client/api/httprouter_registered"
	"easydevops/client/config"
	"easydevops/client/initialization"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"net/http"
)

func init() {
	// 进行项目初始化
	initialization.Initialization()
}

func main() {
	app := httprouter_registered2.RegisterHttprouter()

	clog.Println("http://" + config.Basis.App.Host)
	log.Fatal(http.ListenAndServe(config.Basis.App.Host, app))
}
