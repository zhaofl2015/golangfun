package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

var globalSessions *session.Manager

func init() {
	s := `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`
	cf := new(session.ManagerConfig)
	cf.EnableSetCookie = true
	err := json.Unmarshal([]byte(s), cf)
	if err != nil {
		os.Exit(0)
	}
	globalSessions, _ = session.NewManager("memory", cf)
	go globalSessions.GC()
}

func (c *UserController) Login() {
	c.TplName = "login.html"
}

func (c *UserController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	//	fmt.Println(username, password)
	if username != "admin" || password != "1" {
		c.TplName = "badlogin.html"
		return
	}
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	sess.Set("username", username)
	c.Ctx.Redirect(302, "/")
}

func (c *UserController) Logout() {
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)

	sess.Delete("username")
	c.Ctx.Redirect(302, "/login")
}

func (c *UserController) Fun() {
	fmt.Println("have fun")
	c.Ctx.WriteString("having fun")
}
