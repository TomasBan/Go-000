package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"fmt"
)

/*

Week03 作业题目：
1.基于 errgroup 实现一个http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

*/

func main() {

	egroup := new(errgroup.Group)
	var srv http.Server
	idleConnsClosed := make(chan struct{})

	egroup.Go(func() error {
		log.Println("服务正在运行中...")
		return srv.ListenAndServe()
	})

	egroup.Go(func() error {
		//监听系统信号
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("服务关闭失败: %v", err)
			return err
		}

		fmt.Println("服务已经正常关闭")
		close(idleConnsClosed)
		return nil
	})

	if err := egroup.Wait(); err != nil {
		// Error starting or closing listener:
		log.Printf("errgroup发生错误: %v", err)
	}

	<-idleConnsClosed
}
