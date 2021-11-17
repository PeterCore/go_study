package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloSever)
	fmt.Println("http server start")
	err := srv.ListenAndServe()
	return err
}

func HelloSever(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")

}

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	srv := &http.Server{Addr: ":8888"}

	group.Go(func() error {
		fmt.Println("http")
		go func() {
			<-ctx.Done()
			fmt.Println("http ctx done")
			//srv.Shutdown(context.TODO())
			srv.Shutdown(ctx)
		}()
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		signalChanel := make(chan os.Signal, 1)
		signal.Notify(signalChanel, os.Interrupt, syscall.SIGTERM)
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-signalChanel:
				return nil
			}
		}

	})

	group.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("inject error")
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}
