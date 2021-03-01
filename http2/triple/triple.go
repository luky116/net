package triple

import (
	"bytes"
	"context"
	"net/http"
)

type MsgType uint8
const (
	DataMsgType              = MsgType(1)
	ServerStreamCloseMsgType = MsgType(2)
)

// BufferMsg is the basic transfer unit in one stream
type BufferMsg struct {
	Buffer  *bytes.Buffer
	MsgType MsgType
	Err     error
}

// ProtocolHeader
type ProtocolHeader interface {
	GetStreamID() uint32
	GetPath() string
	FieldToCtx() context.Context
}

type ProtocolHeaderHandler interface {
	ReadFromTripleReqHeader(header *http.Request) ProtocolHeader
	WriteTripleReqHeaderField(header http.Header) http.Header
	WriteTripleFinalRspHeaderField(w http.ResponseWriter)
}


type StreamingRequest struct {
	SendChan chan BufferMsg
	Handler ProtocolHeaderHandler
}

func (sr*StreamingRequest)Read(p []byte)(n int, err error){
	return 0, nil
}

func (sr *StreamingRequest) Close() error{
	return nil
}
