package routes

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type roomoid struct {
	Openid string `json:"openid"`
	// first     string `json:"fisp"`
	// second    string `json:"secp"`
	// third     string `json:"thp"`
	// fouth     string `json:"fop"`
	// fifth     string `json:"fip"`
	// sixth     string `json:"sip"`
	// f_number  int    `json:"fisn"`
	// se_number int    `json:"sen"`
	// th_number int    `json:"thn"`
	// fo_number int    `json:"fon"`
	// fi_number int    `json:"fin"`
	// si_number int    `json:"sin"`
}

func newroom(c *gin.Context) {
	var rid roomoid
	c.BindJSON(&rid)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	roomstr := rid.Openid + timeStr
	fmt.Print(rid.Openid)
	// db := mysql.Getlink()
	// db.Exec("insert into room(roomid,first,sceond,third,fouth,fifth,sixth,f_number,se_number,th_number,fo_number,fi_number,si_number) values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
	// 	roomstr, rid.first, rid.second, rid.third, rid.fouth, rid.fifth, rid.sixth, rid.f_number, rid.se_number, rid.th_number, rid.fo_number, rid.fi_number, rid.si_number)
	c.JSON(200, gin.H{
		"roomid": roomstr,
	})
}
