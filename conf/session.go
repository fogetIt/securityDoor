package conf

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
)


var globalSessions *session.Manager


/*
TODO  not used
 */
func SessionInit() {
	sessionConfig := &session.ManagerConfig{
		Secure: true,
		CookieName: beego.AppConfig.String("SessionName"),
		Domain: beego.AppConfig.DefaultString("SessionDomain", ""),
		ProviderConfig: beego.AppConfig.String("SessionProviderConfig"),
		EnableSetCookie: beego.AppConfig.DefaultBool("SessionAutoSetCookie", true),
		Gclifetime: beego.AppConfig.DefaultInt64("SessionGCMaxLifetime", 60 * 60 * 4),
		Maxlifetime: beego.AppConfig.DefaultInt64("SessionGCMaxLifetime", 60 * 60 * 4),
		CookieLifeTime: beego.AppConfig.DefaultInt("SessionCookieLifeTime", 60 *60 * 24 * 7),
	}
	globalSessions, _ = session.NewManager("redis", sessionConfig)
	go globalSessions.GC()
}
