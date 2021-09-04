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
}

// ProtocolHeader
type ProtocolHeader interface {
	GetPath() string
	FieldToCtx() context.Context
}

// ProtocolHeaderHandler
type ProtocolHeaderHandler interface {
	// ReadFromTripleReqHeader read http header field from http request to ProtocolHeader
	ReadFromTripleReqHeader(header *http.Request) ProtocolHeader

	// WriteTripleReqHeaderField write protocol header fields to http Header
	WriteTripleReqHeaderField(header http.Header) http.Header

	// WriteTripleFinalRspHeaderField write protocol trailer fields to http2 trailer header
	WriteTripleFinalRspHeaderField(w http.ResponseWriter, grpcStatusCode int, grpcMessage string, traceProtoBin int)
}

type StreamingRequest struct {
	SendChan chan BufferMsg
	Handler  ProtocolHeaderHandler
}

func (sr *StreamingRequest) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (sr *StreamingRequest) Close() error {
	return nil
}
