/**
 * @Author: DollarKiller
 * @Description: cmd 配置模块
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:21 2019-09-21
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Node struct {
	FullName string `yaml:"full_name"`
	Branch   string `yaml:"branch"`
}

type base struct {
	App struct {
		DevopsServer string `json:"devops_server" yaml:"devops_server"`
	}
	Devops Node `json:"devops" yaml:"devops"`
}

var (
	Basis *base
)

func init() {
	// 判断配置文件是否存在 如果不存在 则创建
	b, e := easyutils.PathExists("./config.yml")
	if e != nil || b == false {
		createConfig()
		clog.Println("配置文件初始化完成")
		clog.PrintEr("请填写配置文件 并检查是否填写正确")
		os.Exit(0)
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
}

func createConfig() {
	err := ioutil.WriteFile("config.yml", []byte(devposconfig), 00666)
	if err != nil {
		panic("配置文件 创建失败")
	}
}

var devposconfig = `
# devops生成文件   开发机 cli 端

app:
  devops_server: ""        # devops 服务器地址  https:// ...

devops:
  full_name: ""        # 名称 例如 dollarkillerx/easyutils
  branch: "master"     # 分支
`
