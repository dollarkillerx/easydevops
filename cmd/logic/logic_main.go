/**
 * @Author: DollarKiller
 * @Description: 主逻辑
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:31 2019-09-21
 */
package logic

import (
	"easydevops/cmd/config"
	"easydevops/utils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/httplib"
	"os"
	"strings"
)

type Logic struct {
}

// 打包
func (l *Logic) Bale() {
	i := utils.Utils{}

	i.Zip(l.GetName())
	clog.Println("文件打包完毕")
}


// 上传
func (l *Logic) Up() {
	s := config.Basis.App.DevopsServer

	post := httplib.Post(s)
	file := post.PostFile("file", l.GetName()+".zip")
	i, e := file.String()
	if e != nil {
		clog.PrintWa("上传到 devops服务器失败  请检查填写是否正确")
		os.Exit(0)
	}
	if i == "200" {
		clog.Println("打包文件 上传完毕")
	}else {
		clog.PrintWa("打包文件 上传失败")
		os.Exit(0)
	}

	// 文件上传完毕 删除本地文件
	e = os.Remove(l.GetName() + ".zip")
	if e != nil {
		clog.PrintWa("清理冗余数据失败")
	}
}

// 未来会添加读取 .gitignore 来判断哪些文件不更新
func (l *Logic) GetName() string {
	fullName := strings.Replace(config.Basis.Devops.FullName, "/", "1", -1)
	name := fullName + config.Basis.Devops.Branch
	return name
}