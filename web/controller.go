package web

import (
	"net/http"
	"github.com/chandlerxue/proxy"
	"io"
)


type MyHandler int

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := "http://www.baidu.com"
	html := proxy.Fetch(  &url   )
	io.WriteString( res,html )
}

func RunWeb(){
	var h MyHandler
	http.ListenAndServe(":9090", h )
}

