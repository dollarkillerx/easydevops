/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:05 2019-09-16
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Node struct {
	Port               string `yaml:"port"`
	FullName           string `yaml:"full_name"`
	Branch             string `yaml:"branch"`
	Giturl             string `yaml:"giturl"`
	Runname            string `yaml:"runname"`
	Dirpath            string `yaml:"dirpath"`
	Secondarydirectory string `yaml:"secondarydirectory"`
}

type base struct {
	App struct {
		Host       string `yaml:"host"`
		Debug      bool   `yaml:"debug"`
		MaxRequest int    `yaml:"max_request"`
		LogLevel   string `yaml:"log_level"`
		TaskNum    int    `yaml:"task_num"`
	}
	Devops struct {
		Node []Node `yaml:"node"`
	}
}

var (
	Basis *base
)

func init() {

	// 判断配置文件是否存在 如果不存在 则创建
	b, e := easyutils.PathExists("./config.yml")
	if e != nil || b == false {
		createConfig()
		panic("请填写配置文件")
	}

	Basis = &base{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, Basis)
	if e != nil {
		panic(e.Error())
	}

	if Basis.App.Debug {
		log.Println(Basis)
		clog.Println(Basis.Devops)
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
  debug: true
  max_request: 1000
  task_num: 10

# 要自动化部署应用的配置
devops:
  node:
    - port: "8081"                             # 程序运行端口
      full_name: "dollarkillerx/easyutils"     # 名称 例如 dollarkillerx/easyutils
      branch: "master"                         # 分支
      giturl: "https://github"    			   # git pull url 地址 (你要先配置一下秘钥啊!)
      runname: "api"                           # 运行程序的name
      dirpath: "/home/s"                       # 绝对路径
      secondarydirectory: ""                   # 如果有二级目录 就填写在这里
    - port: "8082"                              
      full_name: ""      
      branch: "es"      
      giturl: ""    
      runname: ""    
      dirpath: ""    
      secondarydirectory: ""   
`
