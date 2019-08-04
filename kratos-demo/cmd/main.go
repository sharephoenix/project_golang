package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/log"
	"kratos-demo/internal/server/http"
	"kratos-demo/internal/service"
	"kratos-demo/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//rpc.RpcServer()
	exeDratos()
}

/*启动 rpc 服务*/
func toCaseRpcServer()  {
	rpc.ToCaseServer()
}

/*运行 kratos demo*/
func exeDratos() {
	flag.Parse()

	if err := paladin.Init(); err != nil {
		fmt.Println("errrrrrrrr:::::::::")
		panic(err)
	}
	log.Init(nil) // debug flag: log.dir={path}
	defer log.Close()
	log.Info("kratos-demo start")
	svc := service.New()
	httpSrv := http.New(svc)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
			if err := httpSrv.Shutdown(ctx); err != nil {
				log.Error("httpSrv.Shutdown error(%v)", err)
			}
			log.Info("kratos-demo exit")
			svc.Close()
			cancel()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}