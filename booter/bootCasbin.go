package booter

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	cga "github.com/gohouse/casbin-gorose-adapter"
	"github.com/gohouse/gorose/v2"
)

func BootCasbin(baseConfig *BaseConfig, engin *gorose.Engin) *casbin.Enforcer {
	//enforcer, e := casbin.NewEnforcer("config/casbin_model.conf", cga.NewAdapter(BootGorose()))
	enforcer, err := casbin.NewEnforcer(initCasbinModel(baseConfig), cga.NewAdapter(engin))
	if err!=nil {
		panic(err.Error())
	}
	return enforcer
}

func initCasbinModel(baseConfig *BaseConfig) model.Model {
	var b = baseConfig.CasbinModel
	m := model.NewModel()
	for k,v:=range b{
		m.AddDef(k,k,v)
	}
	//m.AddDef("r", "r", "sub, obj, act")
	//m.AddDef("p", "p", "sub, obj, act")
	//m.AddDef("g", "g", "_, _")
	//m.AddDef("e", "e", "some(where (p.eft == allow))")
	//m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	return m
}
