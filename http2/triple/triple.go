package triple

import "bytes"
import tripleCommon "github.com/dubbogo/triple/pkg/common"

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


type StreamingRequest struct {
	SendChan chan BufferMsg
	Handler tripleCommon.ProtocolHeaderHandler
}

func (sr*StreamingRequest)Read(p []byte)(n int, err error){
	return 0, nil
}

func (sr *StreamingRequest) Close() error{
	return nil
}
