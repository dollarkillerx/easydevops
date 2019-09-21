/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:19 2019-09-21
 */
package main

import (
	"easydevops/server/config"
	"easydevops/server/web/httprouter_registered"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"net/http"
	"os"
)

func init() {
	// 初始化 服务器
	b, e := easyutils.PathExists("./file")
	if e == nil && b == true {
		// 如果存在 就删除历史数据
		e := os.RemoveAll("./file")
		if e != nil {
			clog.PrintWa("历史 文件 删除 失败  请手动删除")
			os.Exit(0)
		}
	}

	// 创建 文件存放目录
	e = easyutils.DirPing("./file")
	if e != nil {
		clog.PrintWa("初始化文件目录失败")
		os.Exit(0)
	}
}


func main() {
	httprouter := httprouter_registered.RegisterHttprouter()

	clog.Println("http://" + config.Basis.App.Host)
	log.Fatal(http.ListenAndServe(config.Basis.App.Host, httprouter))
}
