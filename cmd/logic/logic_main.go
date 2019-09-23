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
	clog.Println("文件开始上传到 devops")
	clog.Println(l.GetName())
	s := config.Basis.App.DevopsServer

	post := httplib.Post(s)
	post.Param("server_key",config.Basis.App.ServerKey)
	post.Param("key",config.Basis.App.Key)
	post.PostFile("file", l.GetName()+".tar.gz")
	i, e := post.String()

	if e != nil {
		clog.PrintWa("上传到 devops服务器失败  请检查填写是否正确")
		os.Exit(0)
	}

	switch i {
	case "200":
		clog.Println("打包文件 上传完毕")
	case "401":
		clog.PrintWa("打包文件 上传失败 401")
		os.Exit(0)
	case "500":
		clog.PrintWa("打包文件 上传失败 500")
		os.Exit(0)
	default:
		clog.PrintWa("未知错误！" + i)
		os.Exit(0)
	}

	// 文件上传完毕 删除本地文件
	e = os.Remove(l.GetName() + ".tar.gz")
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