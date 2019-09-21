/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:03 2019-09-21
 */
package test

import (
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
