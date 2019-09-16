/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:01 2019-09-16
 */
package main

import (
	"easydevops/api/httprouter_registered"
	"easydevops/config"
	"easydevops/initialization"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"net/http"
)

func init() {
	// 进行项目初始化
	initialization.Initialization()
}

func main() {
	app := httprouter_registered.RegisterHttprouter()

	clog.Println("http://" + config.Basis.App.Host)
	log.Fatal(http.ListenAndServe(config.Basis.App.Host, app))
}
