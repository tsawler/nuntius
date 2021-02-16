package connection

import (
	"sync"
	"time"

	log "github.com/golang/glog"
)

// Socket interface to write to the client
type Socket interface {
	WriteJSON(interface{}) error
}

// Connection A user connection
type Connection struct {
	sync.Mutex

	SocketID  string
	Socket    Socket
	CreatedAt time.Time
}

// New Creates a new Subscriber
func New(socketID string, s Socket) *Connection {
	log.Infof("Creating a new Subscriber %+v", socketID)

	return &Connection{SocketID: socketID, Socket: s, CreatedAt: time.Now()}
}

// Publish the message to websocket attached to this client
func (conn *Connection) Publish(m interface{}) {
	conn.Lock()
	defer conn.Unlock()

	if err := conn.Socket.WriteJSON(m); err != nil {
		log.Errorf("error writing json into Socket, %+v", err)
	}
}
