package sdk

type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}
