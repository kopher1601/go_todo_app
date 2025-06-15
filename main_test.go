package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	// HTTP 서버의 반환값을 검증한다
	want := fmt.Sprintf("Hello, %s", in)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, string(got))
	}

	// 종료 알림을 보낸다.
	cancel()
	// run 함수의 반환값을 검증한다
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
