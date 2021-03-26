package triple

import(
	"io"
	"net/http"
)

type ResponseBody struct {
	rawHttp2ResponseBody io.ReadCloser
	trailerChan chan http.Header
}


func (rsp * ResponseBody)Read(p []byte) (n int, err error) {
	return rsp.rawHttp2ResponseBody.Read(p)
}

func (rsp*ResponseBody) Close() error{
	return rsp.rawHttp2ResponseBody.Close()
}

func (rsp*ResponseBody) GetTrailer()http.Header{
	trailer := <- rsp.trailerChan
	return trailer
}
func (rsp*ResponseBody) GetTrailerChan() chan http.Header{
	return rsp.trailerChan
}

func (rsp*ResponseBody) SetTrailerChan(trailerChan chan http.Header){
	rsp.trailerChan = trailerChan
}

func (rsp *ResponseBody) SetRawHttp2ResponseBody(readCloser io.ReadCloser){
	rsp.rawHttp2ResponseBody = readCloser
}