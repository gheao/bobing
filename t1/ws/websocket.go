package ws

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// SystemMessage 系统消息
	SystemMessage = iota
	// BroadcastMessage 广播消息(正常的消息)
	BroadcastMessage
	// HeartBeatMessage 心跳消息
	HeartBeatMessage
	// ConnectedMessage 上线通知
	ConnectedMessage
	// DisconnectedMessage 下线通知
	DisconnectedMessage
)

type AliveList struct {
	ConnList     map[string]*Client
	register     chan *Client
	destroy      chan *Client
	broadcast    chan Message
	allBroadcast chan Message
	//cancel       chan int
	Len int
}

// Client socket客户端
type Client struct {
	ID   string
	conn *websocket.Conn
	//cancel chan int
}
type Message struct {
	ID      string
	Content string
	SentAt  int64
	Type    int // <- SystemMessage 等类型就是这里了
}

// 向某个链接发送消息
func (al *AliveList) sendMessage(id string, msg Message) {

	data, _ := json.Marshal(msg)
	fmt.Print("\nsend\n")

	al.ConnList[id].conn.WriteMessage(websocket.TextMessage, data)
}

// 服务器系统广播
func (al *AliveList) SysBroadcast(code int, msg Message) {
	for id := range al.ConnList {
		al.sendMessage(id, msg)
	}
}

func (al *AliveList) Listen() {
	for {
		select {
		// 链接注册
		case client := <-al.register:
			al.ConnList[client.ID] = client
			al.Len++
			al.SysBroadcast(ConnectedMessage, Message{
				ID:      client.ID,
				Content: "connected",
				SentAt:  time.Now().Unix(),
				Type:    ConnectedMessage,
			})
		// 链接注销
		case client := <-al.destroy:
			_ = client.conn.Close()
			delete(al.ConnList, client.ID)
			al.Len--
		// 消息除源广播
		case msg := <-al.broadcast:
			for id := range al.ConnList {
				if id != msg.ID {
					al.sendMessage(id, msg)
				}
			}
		// 消息广播
		case msg := <-al.allBroadcast:
			for id := range al.ConnList {
				al.sendMessage(id, msg)
			}
		}
	}
}

func (al *AliveList) Register(c *Client) {
	fmt.Print("\nregister\n")
	al.register <- c
	fmt.Print("\nregister\n")
}

func (al *AliveList) Destroy(c *Client) {
	al.destroy <- c
}

func (al *AliveList) Broadcast(msg Message) {
	al.broadcast <- msg
}

func (al *AliveList) AllBroadcast(msg Message) {
	al.allBroadcast <- msg
}

func ClientInit(conn *websocket.Conn, id string) *Client {
	c := new(Client)
	c.conn = conn
	c.ID = id
	return c
}
