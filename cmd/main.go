/**
 * @Author: DollarKiller
 * @Description: 自动更新工具 开发机用cli
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 09:58 2019-09-21
 */
package main

import (
	"easydevops/cmd/logic"
	"easydevops/utils"
	"github.com/dollarkillerx/easyutils/clog"
)

// zip打包当前目录

func init() {
	clog.Println("系统初始化 完毕")
}

func main() {
	loc := logic.Logic{}

	loc.Bale()  // 打包
	loc.Up()    // 上传

	util := utils.Utils{}
	util.Github()  // 同步github

	clog.Println("更新以及全部提交")
}
