/**
 * @Author: DollarKiller
 * @Description: 中心配置文件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:42 2019-09-21
 */
package controller

import (
	"easydevops/common"
	"easydevops/server/config"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/gcache"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func UpFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseMultipartForm(200 << 20)  // 设置上传文件最大大小  200M
	r.ParseForm()
	resp := common.Resp{}
	server_key := r.FormValue("server_key")

	if server_key != config.Basis.App.ServerKey {
		clog.Println(server_key)
		resp.Auth401(w)
		return
	}
	key := r.FormValue("key")
	file, header, e := r.FormFile("file")
	if file != nil {
		defer file.Close()
	}
	if strings.Index(header.Filename,"tar.gz") == -1 {
		resp.Auth401(w)
		return
	}
	key = easyutils.Sha256Encode(key)

	// 判断key是否存在,如果存在 就删除原来的数据
	get, b := gcache.CacheGet(key)
	if b {
		// 删除原来的数据
		os.Remove(get.(string))
	}

	fileh := easyutils.SuperRand()
	filepath := "./file/" + fileh + ".tar.gz"
	e = gcache.CacheSet(key, filepath)
	if e != nil {
		resp.Bad500(w)
		return
	}

	bytes, e := ioutil.ReadAll(file)
	if e != nil {
		resp.Bad500(w)
		return
	}

	e = ioutil.WriteFile(filepath, bytes, 00666)
	if e != nil {
		resp.Bad500(w)
		return
	}

	resp.Ok200(w)
}

func GetFile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	r.ParseForm()
	server_key := r.FormValue("server_key")
	resp := common.Resp{}

	if server_key != config.Basis.App.ServerKey {
		resp.Auth401(w)
		return
	}
	key := r.FormValue("key")
	key = easyutils.Sha256Encode(key)
	get, b := gcache.CacheGet(key)
	if !b {
		resp.Auth401(w)
		return
	}

	filepath := get.(string)
	bytes, e := ioutil.ReadFile(filepath)
	if e != nil {
		resp.Auth401(w)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type","application/octet-stream")
	w.Header().Set("Content-Disposition","attachment; filename=hc.tar.gz")
	w.Write(bytes)
	// 如果文件 下载完毕就删除文件
	os.Remove(filepath)
}
