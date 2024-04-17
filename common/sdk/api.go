package sdk

import "sync"

const (
	MsgTypeText      = "text"
	MsgTypeAck       = "ack"
	MsgTypeReConn    = "reConn"
	MsgTypeHeartbeat = "heartbeat"
	MsgLogin         = "loginMsg"
)

type Chat struct {
	NickName         string
	UserId           string
	SessionId        string
	connect          *connect
	closeChan        chan struct{}
	MsgClientIDTable map[string]uint64
	sync.RWMutex
}

type Message struct {
	Type       string
	Name       string
	FormUserID string
	ToUserID   string
	Content    string
	Session    string
}

func NewChat(serverAddr, nackName, userId, SessionId string) *Chat {
	return &Chat{
		NickName:  nackName,
		UserId:    userId,
		SessionId: SessionId,
		connect:   newConnect(serverAddr),
	}
}

func (c *Chat) GetCurClientID() {

}

func (c *Chat) Send(msg *Message) {
	c.connect.Send(msg)
}

func (c *Chat) Recv() <-chan *Message {
	return c.connect.recvChan
}

func (c *Chat) Close() {

}
