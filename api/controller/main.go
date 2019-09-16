/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:08 2019-09-16
 */
package controller

import (
	"easydevops/config"
	"easydevops/defs"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/formband"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func reLanuch(sh string) {
	cmd := exec.Command("sh", sh)
	err := cmd.Start()
	if err != nil {
		clog.Println(err)
	}
	err = cmd.Wait()
}

func Rundevops(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	clog.Println("收到github更新任务 进行更新")

	data := defs.GithubAPI{}
	err := formband.Band(r, &data)
	if err != nil {
		clog.Println(err)
		return
	} else {
		log.Println(data)
	}
	shpath := checkexis(data.Repository.FullName, data.Branch)
	if shpath != "" {
		reLanuch(shpath)
	}
}

// 查询是否存在这个  如果找到 返回对应的脚本地址
func checkexis(name, branch string) string {

	for _, k := range config.Basis.Devops.Node {
		// 如果找到了
		if k.FullName == name {
			// 再判断brench
			if strings.Index(branch, k.Branch) != -1 {
				// 返回脚本地址
				name := strings.Replace(k.FullName, "/", "1", -1)
				shName := name + k.Port + k.Branch + ".sh"
				shPath := "./sh/" + shName
				return shPath
			}
		}
	}

	return ""
}
