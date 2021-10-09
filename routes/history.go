package routes

import (
	"fmt"
	"t1/dao/mysql"

	"github.com/gin-gonic/gin"
)

type history struct {
	id  string
	sum int
	a   int
	b   int
	c   int
	d   int
	e   int
	f   int
}
type opid struct {
	Openid string `json:"openid"`
}

func getHistory(c *gin.Context) {
	var openid opid
	c.BindJSON(&openid)
	db := mysql.Getlink()
	fmt.Print(openid)
	serch, err := db.Query("select id from userinfo where id = ?", openid.Openid)
	if err != nil {
		fmt.Print("err")
	}
	var his history
	serch.Scan(&his.id, &his.sum, &his.a, &his.b, &his.c, &his.d, &his.e, &his.f)
	fmt.Print(openid.Openid)
	c.JSON(200, gin.H{
		"sum": his.sum,
		"a":   his.a,
		"b":   his.b,
		"c":   his.c,
		"d":   his.d,
		"e":   his.e,
		"f":   his.f,
	})
}
