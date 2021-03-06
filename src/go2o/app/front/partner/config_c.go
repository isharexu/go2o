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
	"encoding/json"
	"github.com/atnet/gof"
	"github.com/atnet/gof/app"
	"github.com/atnet/gof/web"
	"go2o/core/domain/interface/partner"
	"go2o/core/service/dps"
	"html/template"
	"net/http"
	"time"
)

type configC struct {
	app.Context
}

//资料配置
func (this *configC) Profile(w http.ResponseWriter, r *http.Request, partnerId int) {
	p, _ := dps.PartnerService.GetPartner(partnerId)
	p.Pwd = ""
	p.Secret = ""
	p.ExpiresTime = time.Now().Unix()

	js, _ := json.Marshal(p)

	this.Context.Template().Execute(w,
		func(m *map[string]interface{}) {
			(*m)["entity"] = template.JS(js)
		},
		"views/partner/conf/profile.html")
}

func (this *configC) Profile_post(w http.ResponseWriter, r *http.Request, partnerId int) {
	var result gof.JsonResult
	r.ParseForm()

	e := partner.ValuePartner{}
	web.ParseFormToEntity(r.Form, &e)

	//更新
	origin, _ := dps.PartnerService.GetPartner(partnerId)
	e.ExpiresTime = origin.ExpiresTime
	e.Secret = origin.Secret
	e.JoinTime = origin.JoinTime
	e.LastLoginTime = origin.LastLoginTime
	e.LoginTime = origin.LoginTime
	e.Pwd = origin.Pwd
	e.UpdateTime = time.Now().Unix()
	e.Id = partnerId

	id, err := dps.PartnerService.SavePartner(partnerId, &e)

	if err != nil {
		result = gof.JsonResult{Result: false, Message: err.Error()}
	} else {
		result = gof.JsonResult{Result: true, Message: "", Data: id}
	}
	w.Write(result.Marshal())
}

//站点配置
func (this *configC) SiteConf(w http.ResponseWriter, r *http.Request, partnerId int) {
	conf := dps.PartnerService.GetSiteConf(partnerId)
	js, _ := json.Marshal(conf)

	this.Context.Template().Execute(w,
		func(m *map[string]interface{}) {
			(*m)["entity"] = template.JS(js)
		},
		"views/partner/conf/site_conf.html")
}

func (this *configC) SiteConf_post(w http.ResponseWriter, r *http.Request, partnerId int) {
	var result gof.JsonResult
	r.ParseForm()

	e := partner.SiteConf{}
	web.ParseFormToEntity(r.Form, &e)

	//更新
	origin := dps.PartnerService.GetSiteConf(partnerId)
	e.Host = origin.Host
	e.PartnerId = partnerId

	err := dps.PartnerService.SaveSiteConf(partnerId, &e)

	if err != nil {
		result = gof.JsonResult{Result: false, Message: err.Error()}
	} else {
		result = gof.JsonResult{Result: true, Message: ""}
	}
	w.Write(result.Marshal())
}

//销售配置
func (this *configC) SaleConf(w http.ResponseWriter, r *http.Request, partnerId int) {
	conf := dps.PartnerService.GetSaleConf(partnerId)
	js, _ := json.Marshal(conf)

	this.Context.Template().Execute(w,
		func(m *map[string]interface{}) {
			(*m)["entity"] = template.JS(js)
		},
		"views/partner/conf/sale_conf.html")
}

func (this *configC) SaleConf_post(w http.ResponseWriter, r *http.Request, partnerId int) {
	var result gof.JsonResult
	r.ParseForm()

	e := partner.SaleConf{}
	web.ParseFormToEntity(r.Form, &e)

	e.PartnerId = partnerId

	err := dps.PartnerService.SaveSaleConf(partnerId, &e)

	if err != nil {
		result = gof.JsonResult{Result: false, Message: err.Error()}
	} else {
		result = gof.JsonResult{Result: true, Message: ""}
	}
	w.Write(result.Marshal())
}
