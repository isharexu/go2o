/**
 * Copyright 2014 @ ops Inc.
 * name :
 * author : newmin
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package partner

import (
	"fmt"
	"github.com/atnet/gof/app"
	"github.com/atnet/gof/web"
	"go2o/app/front"
	"go2o/app/session"
	"net/http"
)

type mainC struct {
	app.Context
	*front.WebCgi
}

//入口
func (this *mainC) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<script>location.replace('/dashboard')</script>"))
}

func (this *mainC) Logout(w http.ResponseWriter, r *http.Request) {
	session.LSession.PartnerLogout(w, r)
	w.Write([]byte("<script>location.replace('/login')</script>"))
}

//商户首页
func (this *mainC) Dashboard(w http.ResponseWriter, r *http.Request) {
	pt, err := session.LSession.GetCurrentSessionFromCookie(r)
	if err != nil {
		w.Write([]byte("<script>location.replace('/login')</script>"))
		return
	}

	var mf web.TemplateMapFunc = func(m *map[string]interface{}) {
		(*m)["partner"] = pt
		(*m)["loginIp"] = r.Header.Get("USER_ADDRESS")
	}
	this.Context.Template().Render(w, "views/partner/dashboard.html", mf)
}

//商户汇总页
func (this *mainC) Summary(w http.ResponseWriter, r *http.Request) {
	pt, err := session.LSession.GetCurrentSessionFromCookie(r)
	if err != nil {
		return
	}
	this.Context.Template().Render(w,
		"views/partner/summary.html",
		func(m *map[string]interface{}) {
			(*m)["partner"] = pt
			(*m)["loginIp"] = r.Header.Get("USER_ADDRESS")
		})
}

func (this *mainC) Upload_post(w http.ResponseWriter, r *http.Request) {
	ptid, _ := session.LSession.GetPartnerIdFromCookie(r)
	r.ParseMultipartForm(20 * 1024 * 1024 * 1024) //20M
	for f, _ := range r.MultipartForm.File {
		w.Write(this.WebCgi.Upload(f, w, r, fmt.Sprintf("%d/item_pic/", ptid)))
		break
	}
}
