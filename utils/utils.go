/**
 * @Author: DollarKiller
 * @Description:  本项目 通用utils
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:01 2019-09-21
 */
package utils

import (
	"bytes"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/compression"
	"os/exec"
)

type Utils struct {
}

func (u *Utils) Exec(sh string, arg ...string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(sh, arg...)
	cmd.Stdout = &stdout // 输出
	cmd.Stderr = &stderr // 输出错误
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func (u *Utils) Zip(name string) {
	zip := compression.Tar{}
	err := zip.Tar(".", name)
	if err != nil {
		panic(err)
	}
}

func (u *Utils) Github() {
	e, s, i2 := u.Exec("git", "add", ".")
	if e != nil {
		clog.Println(s)
		clog.Println(i2)
	}
	clog.Println("github add 完毕")
	e, s, i2 = u.Exec("git", "commit", "-m", "EasyDevOps")
	if e != nil {
		clog.Println(s)
		clog.Println(i2)
	}
	clog.Println("github commit 完毕")
	e, s, i2 = u.Exec("git", "push")
	if e != nil {
		clog.Println(s)
		clog.Println(i2)
	}
	clog.Println("github push 完毕")
}
