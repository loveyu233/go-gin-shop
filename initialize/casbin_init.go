package initialize

import (
	"github.com/casbin/casbin"
)

var enforcer *casbin.Enforcer

func InitCasbin() {
	enforcer = casbin.NewEnforcer("./static/casbin/model.conf", "./static/casbin/policy.csv")
}

func Check(sub, obj, act string) bool {
	ok := enforcer.Enforce(sub, obj, act)
	return ok
}
