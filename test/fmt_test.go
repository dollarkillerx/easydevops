/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 11:57 2019-09-16
 */
package test

import (
	"fmt"
	"github.com/dollarkillerx/easyutils/clog"
	"testing"
)

func TestFmt(t *testing.T) {
	sprintf := fmt.Sprintf(sh1, "ss", "sd", "cs", "ss")
	clog.Println(sprintf)
}

var sh1 = `
#! /bin/sh

kill -9 $(grep %s)

cd %s

git pull %s


./%s &
`