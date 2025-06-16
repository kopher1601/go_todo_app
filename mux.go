package main

import (
	"net/http"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// 정적 분석 오류를 회피하기 위해 명시적으로 반환값을 버린다.
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	return mux
}
