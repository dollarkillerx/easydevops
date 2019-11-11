/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:03 2019-09-21
 */
package test

import (
	"easydevops/utils"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	as := []string{"sadas", "regre"}
	a(as...)
}

func a(arg ...string) {
	for _, k := range arg {
		log.Println(k)
	}
}


func TestRun(t *testing.T) {
	//lsof -i :8081 | awk '{print $2}'> tmp | pid=$(awk 'NR==2{print}' tmp) | echo $pid;
	i := utils.Utils{}

	e, s, i2 := i.Exec( "sh","lsof", "-i", ":8081", "| awk '{print $2}'> tmp | pid=$(awk 'NR==2{print}' tmp) | echo $pid;")
	clog.PrintWa(e)
	clog.Println(s)
	clog.Println(i2)
}