package webJsonServe

import (
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"golang.org/x/net/websocket"
	"log"
	"sync"
	"unsafe"
)

type WS struct {
	Conn  *websocket.Conn
	Order *chan int
}

type Socket struct {
	websocket.Handler
	Kind    string
	mutex   sync.RWMutex
	clients map[string]WS
}

func NewSocket(kind string) *Socket {
	s := new(Socket)
	s.Kind = kind
	s.clients = make(map[string]WS)
	s.Handler = websocket.Handler(s.PushHandler)
	return s
}

// 添加客户端至ClientsList
func (s *Socket) pushHandler(kind string, ws *websocket.Conn) {
	o := make(chan int)
	log.Printf(kind + ws.RemoteAddr().String() + fmt.Sprintf("_%x", unsafe.Pointer(ws)))
	s.mutex.Lock()
	s.clients[kind+ws.RemoteAddr().String()+fmt.Sprintf("_%x", unsafe.Pointer(ws))] = WS{Conn: ws, Order: &o}
	s.mutex.Unlock()
	for {
		order := <-o
		switch order {
		case 0:
			return
		}
	}
}
func (s *Socket) PushHandler(ws *websocket.Conn) {
	s.pushHandler(s.Kind, ws)
}
func (s *Socket) Push(v interface{}) {
	s.push(s.Kind, v)
}
func (s *Socket) push(kind string, v interface{}) {
	var err error
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for key, ws := range s.clients {
		go func(key string, ws WS) {
			if key[0:len(kind)] != kind {
				return
			} //判断是否是相同类型的监听
			err = websocket.JSON.Send(ws.Conn, v)
			if err != nil {
				beego.BeeLogger.Debug(key + " " + err.Error())
				//log.Printf(key + " " + err.Error())
				s.removeClient(key)
			}
		}(key, ws)
	}
}
func (s *Socket) removeClient(key string) {
	s.mutex.Lock()
	ch := s.clients[key].Order
	delete(s.clients, key)
	s.mutex.Unlock()
	if ch != nil {
		*ch <- 0
	}
}
