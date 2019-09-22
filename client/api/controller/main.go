/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:08 2019-09-16
 */
package controller

import (
	"easydevops/client/config"
	"easydevops/client/defs"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/compression"
	"github.com/dollarkillerx/easyutils/formband"
	"github.com/dollarkillerx/easyutils/gemail"
	"github.com/dollarkillerx/easyutils/httplib"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
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
		//log.Println(data)
		clog.Println("接受到一个push请求")
	}

	// 更具commit 判断 是否是真的更新
	if data.Commits[0].Message != "EasyDevOps" {
		// 直接返回
		return
	}

	shpath := checkexis(data.Repository.FullName, data.Branch)
	if shpath != "" {
		reLanuch(shpath)
	}
}

// 下载对应脚本
func dow(name, branch string) error {
	for _,k := range config.Basis.Devops.Node {
		if k.FullName == name {
			if strings.Index(branch, k.Branch) != -1 {
				// 找到对应下载 脚本的地址
				post := httplib.Post(config.Basis.App.DevopsServer)
				post.Param("server_key",config.Basis.App.ServerKey)
				post.Param("key",k.Key)
				response, e := post.Response()
				defer response.Body.Close()
				if e != nil {
					if k.Email != "" {
						gemail.SendNifoLog([]string{k.Email},"EasyDevops",k.FullName + "  " + k.Branch + "   更新失败")
					}
				}
				if response.StatusCode != 200 {
					if k.Email != "" {
						gemail.SendNifoLog([]string{k.Email},"EasyDevops",k.FullName + "  " + k.Branch + "   更新失败")
					}
				}

				// 文件 下载到对应目录

				filepath := k.Dirpath + "/easydevops"
				// 清理原始数据
				b, e := easyutils.PathExists(filepath)
				if e != nil || b == false {
					os.RemoveAll(filepath)
				}
				e = easyutils.DirPing(filepath)
				if e!= nil {
					clog.Println(e)
				}

				bytes, e := ioutil.ReadAll(response.Body)
				if e != nil {
					if k.Email != "" {
						gemail.SendNifoLog([]string{k.Email},"EasyDevops",k.FullName + "  " + k.Branch + "   更新失败")
					}
				}
				// 下载文件到指定位置
				e = ioutil.WriteFile("easydevops.tar.gz", bytes, 00666)
				if e != nil {
					if k.Email != "" {
						gemail.SendNifoLog([]string{k.Email},"EasyDevops",k.FullName + "  " + k.Branch + "   更新失败")
					}
				}

				// 解压文件
				tar := compression.Tar{}
				e = tar.UnTar(filepath+"/easydevops.tar.gz", filepath)
				if e != nil {
					if k.Email != "" {
						gemail.SendNifoLog([]string{k.Email},"EasyDevops",k.FullName + "  " + k.Branch + "   更新失败")
					}
				}
			}
		}
	}
	return errors.New("not dev")
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

