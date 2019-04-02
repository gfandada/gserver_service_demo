package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/gfandada/gserver"
)

func StartAdminHttp(webaddr string) error {
	adminServeMux := http.NewServeMux()
	adminServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	adminServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	adminServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	adminServeMux.HandleFunc("/debug/pprof/heap", pprof.Profile)
	adminServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	err := http.ListenAndServe(webaddr, adminServeMux)
	if err != nil {
		x := fmt.Sprintf("http.ListenAdServe(\"%s\") failed (%s)", webaddr, err.Error())
		fmt.Println(x)
		return err
	}
	return nil
}

func pproftest() {
	go func() {
		if err := StartAdminHttp("127.0.0.1:7082"); err != nil {
			os.Exit(-1)
		}
	}()
}

func main() {
	pproftest()
	logger := "C:/Users/Administrator/go/src/gserver_service_demo/cfg/gamelogger.xml"
	game := "C:/Users/Administrator/go/src/gserver_service_demo/cfg/game.json"
	gserver.RunService(logger, game, RegisterServices())
}
