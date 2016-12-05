package main

import (
	"github.com/songtianyi/rrframework/logs"
	"net/http"
)

var (
	wslog *logs.BeeLogger
)

func main() {
	wslog = logs.NewLogger(0)
	wslog.SetLogger("websocket", `{"level": 7, "channelSize": 1000}`)
	http.Handle("/", http.FileServer(http.Dir(".")))
	wslog.Info("hello world!")
	wslog.Debug("hehe da!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
