/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:05 2019-09-16
 */
package config

import (
	"easydevops/utils"
	"github.com/dollarkillerx/easyutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Node struct {
	Port               string `yaml:"port"`
	FullName           string `yaml:"full_name"`
	Key                string `json:"key" yaml:"key"`
	Branch             string `yaml:"branch"`
	Runname            string `yaml:"runname"`
	Dirpath            string `yaml:"dirpath"`
	Secondarydirectory string `yaml:"secondarydirectory"`
	Email              string `yaml:"email"`
}

type base struct {
	App struct {
		Host         string `yaml:"host"`
		DevopsServer string `json:"devops_server" yaml:"devops_server"`
		ServerKey    string `json:"server_key" yaml:"server_key"`
	}
	Devops struct {
		Node []Node `yaml:"node"`
	}
}

var (
	Basis *base
)

func init() {
	hasCode := ""
	// 写入热加载
	for {
		select {
		case <-time.After(time.Millisecond * 200):
			md5, e := utils.Util().Md5File("./config.yml")
			if e != nil {
				// 判断配置文件是否存在 如果不存在 则创建
				b, e := easyutils.PathExists("./config.yml")
				if e != nil || b == false {
					createConfig()
					log.Fatalln("请填写配置文件")
				}
			} else {
				if md5 != hasCode {
					hasCode = md5
					// 重载配置文件
					log.Println("重载配置文件")
					initConf()
				}
			}
		}
	}

}

// 载入conf
func initConf() {
	Basis = &base{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, Basis)
	if e != nil {
		panic(e.Error())
	}
}

func createConfig() {
	err := ioutil.WriteFile("config.yml", []byte(devposconfig), 00666)
	if err != nil {
		panic("配置文件 创建失败")
	}
}

var devposconfig = `
# devops生成文件    请填写完成检查确认后 再 运行啊   (每次运行会重写sh模块)


# 这devops系统配置
app:
  host: "0.0.0.0:8083"
  devops_server: "http://0.0.0.0:8083/getfile"        # devops 服务器地址  https:// ...
  server_key: "xxxxxx"                               # 上传 devops 服务器秘钥

# 要自动化部署应用的配置
devops:
  node:
    - port: "8081"                             # 程序运行端口
      full_name: "dollarkillerx/easyutils"     # 名称 例如 dollarkillerx/easyutils
      branch: "master"                         # 分支
      key: "key"                               # 文件同步 秘钥  (服务器同步数据时需要)
      runname: "api"                           # 运行程序的name
      dirpath: "/home/s"                       # 绝对路径
      secondarydirectory: ""                   # 如果有二级目录 就填写在这里
      email: ""                                # 部署成功收邮件通知 
    - port: "8082"
      full_name: ""
      branch: "es"
      key: "key"
      runname: ""
      dirpath: ""
      secondarydirectory: ""
      email: ""
`
