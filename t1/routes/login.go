package routes

import (
	"fmt"
	"t1/dao/mysql"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

func userLogin(c *gin.Context) {
	// 用户登录code
	code := c.Query("code")
	// 用户小程序id
	id := `wxb27cb3df6158fc0e`
	// 小程序secret
	secret := `001a15eeb9b2e60acfb21ce896e3885f`
	// grant_type const
	grant_type := `authorization_code`
	// 请求路径

	resp, err := http.Get(`https://api.weixin.qq.com/sns/jscode2session?appid=` + id + `&secret=` + secret + `&js_code=` + code + `&grant_type=` + grant_type)

	if err != nil {
		fmt.Print("err")
	}
	//fmt.Print(resp)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var res result
	_ = json.Unmarshal(body, &res)
	fmt.Printf("%#v", res)
	fmt.Print("\n")
	//用openid构建数据库
	// var openid opid
	// c.BindJSON(&openid)

	db := mysql.Getlink()
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	db.Exec("insert into userinfo(id,sum,a,b,c,d,e,f) values(?,?,?,?,?,?,?,?)", res.Openid, 0, 0, 0, 0, 0, 0, 0)
	defer db.Close()
}
