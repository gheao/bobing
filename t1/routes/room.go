package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"t1/dao/mysql"
	"t1/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var al = new(ws.AliveList)
var msg ws.Message

type contact struct {
	Roomid    string `json:"roomid"`
	Type      int    `json:"type"`
	Players   []user `json:"players"`
	Index     int    `json:"index"`
	Touzi     [6]int `json:"touzi"`
	Showindex int    `json:"showIndex"`
	Next      int    `json:"next"`
	Openid    string `json:"openid"`
	Result    int    `json:"result"`
	Userinfo  info   `json:"userinfo"`
	Ismaster  bool   `json:"ismaster"`
}
type info struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Language  string `json:"language"`
}
type user struct {
	Identify string `json:"identify"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

func checkOrigin(r *http.Request) bool {
	return true
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func cmdWebSocket(c *gin.Context) {

	wsConn, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	var client *ws.Client
	defer wsConn.Close()
	var contxt contact
	db := mysql.Getlink()
	var t int

	msgtype := make(chan int)

	for {
		_, message, _ := wsConn.ReadMessage()
		fmt.Println(message)
	}

	for {
		t = <-msgtype
		if t == 0 {
			var len int
			serch, _ := db.Query("select max(number) from  roomuser where roomid= ?", contxt.Roomid)
			serch.Scan(&len)
			for i := 0; i <= len; i++ {

				var n, id, r, h string
				var num int
				serch.Scan(&n, &id, &r, &num, &h)
				contxt.Players[i].Name = n
				contxt.Players[i].Identify = h
				contxt.Players[i].Avatar = ""
			}
			co, _ := json.Marshal(contxt)
			str := string(co)
			msg.Content = str
			al.AllBroadcast(msg)
		} else if t == 1 {
			var len int
			serch1, _ := db.Query("select max(number) from  roomuser where roomid= ?", contxt.Roomid)
			serch1.Scan(&len)
			serch1, _ = db.Query("select * from roomuser where userid=?", contxt.Openid)
			var n, id, r, h string
			var num int
			serch1.Scan(&n, &id, &r, &num, &h)
			contxt.Index = num
			contxt.Next = (num + 1) % len
			serch1, _ = db.Query("select userid from roomuser where roomid=? and number=?", contxt.Roomid, contxt.Next)
			serch1.Scan(&n)
			contxt.Openid = n
			co, _ := json.Marshal(contxt)
			str := string(co)
			msg.Content = str
			al.AllBroadcast(msg)
		} else if t == 2 {
			contxt.Index = 0
			co, _ := json.Marshal(contxt)
			str := string(co)
			fmt.Println(str)
			msg.Content = str
			al.AllBroadcast(msg)
		}

	}
}
