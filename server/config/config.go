/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:18 2019-09-21
 */
package config

import (
	"github.com/dollarkillerx/easyutils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type base struct {
	App struct {
		Host       string `yaml:"host"`
		ServerKey string `json:"server_key" yaml:"server_key"`
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
}

func createConfig() {
	err := ioutil.WriteFile("config.yml", []byte(devposconfig), 00666)
	if err != nil {
		panic("配置文件 创建失败")
	}
}

var devposconfig = `
# easydevops 自动生成文件   server 服务器端

app:
  host: "0.0.0.0:8083"  # 服务器运行
  server_key: ""        # 服务器key   用于用户验证
`

