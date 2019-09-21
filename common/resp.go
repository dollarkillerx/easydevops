/**
 * @Author: DollarKiller
 * @Description: 通用返回
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:51 2019-09-21
 */
package common

import "net/http"

type Resp struct {

}

func (r *Resp) Auth401(w http.ResponseWriter) {
	w.WriteHeader(401)
	w.Write([]byte("401"))
}

func (r *Resp) Ok200(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Write([]byte("200"))
}

func (r *Resp) Bad500(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Write([]byte("500"))
}