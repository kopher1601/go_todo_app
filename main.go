package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with %s", url)

	s := &http.Server{
		// 인수로 받은 net.Listener 를 이용하므로 Addr 필드는 지정하지 않는다.
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)
	// 다른 고루틴에서 HTTP 서버를 실행한다.
	eg.Go(func() error {
		// http.ErrServerClosed 는
		// http.Server.Shutdown() 가 정상 종료된 것을 나타내므로 이상 처리가 아니다.
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %v", err)
			return err
		}
		return nil
	})

	// 채널로부터의 알림(종료 알림)을 기다린다.
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// Go 메서드로 실행한 다른 고루틴의 종료를 기다린다.
	return eg.Wait()
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
