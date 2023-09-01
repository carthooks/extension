package extension

import (
	"bufio"
	"bytes"
	"context"
	"net"
	"net/http"

	pb "github.com/carthooks/extension/service"
	"github.com/gin-gonic/gin"
)

func (s Server) OnHttpRequest(c context.Context, r *pb.HttpRequest) (rsp *pb.HttpResponse, err error) {
	ginResponse := ResponseWriter{
		response: &pb.HttpResponse{},
		header:   http.Header{},
	}

	ginContext := gin.CreateTestContextOnly(ginResponse, s.HttpGinEngine)
	req, err := http.NewRequest(r.Method, r.Uri, bytes.NewReader(r.Body))
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	for k, v := range r.Headers {
		header.Set(k, v)
	}
	req.Header = header
	ginContext.Request = req

	s.HttpGinEngine.HandleContext(ginContext)
	ginResponse.response.Headers = map[string]string{}
	for k, v := range ginResponse.Header() {
		for _, vv := range v {
			ginResponse.response.Headers[k] = vv
		}
	}
	return ginResponse.response, nil
}

type ResponseWriter struct {
	response *pb.HttpResponse
	header   http.Header
}

func (w ResponseWriter) Header() http.Header {
	return w.header
}

func (w ResponseWriter) WriteHeader(statusCode int) {
	w.response.Status = int32(statusCode)
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.response.Body = append(w.response.Body, b...)
	return len(b), nil
}

func (w ResponseWriter) WriteString(s string) (int, error) {
	w.response.Body = append(w.response.Body, []byte(s)...)
	return len(s), nil
}

func (w ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return nil, nil, nil
}

func (w ResponseWriter) Flush() {
}

func (w ResponseWriter) CloseNotify() <-chan bool {
	return nil
}

func (w ResponseWriter) Status() int {
	return int(w.response.Status)
}

func (w ResponseWriter) Size() int {
	return len(w.response.Body)
}

func (w ResponseWriter) Written() bool {
	return false
}

func (w ResponseWriter) WriteHeaderNow() {
}

func (w ResponseWriter) Pusher() http.Pusher {
	return nil
}
