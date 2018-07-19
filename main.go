package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) <= 1{
        log.Fatal("port number missing!")
    }

    port := os.Args[1]
    addressBind := "127.0.0.1"
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    
    proxy.OnRequest().DoFunc(
        //You can test on http://httpbin.org/headers
        func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
            r.Header.Set("X-GoProxy","yxorPoG-X from "+addressBind+":"+port)
            log.Print("Adding X-GoProxy to request")
            return r,nil
        })

    log.Print("HTTP Proxy running on http://"+addressBind+":",port)
    log.Fatal(http.ListenAndServe(addressBind+":"+port, proxy))
}


