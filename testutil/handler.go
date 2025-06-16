package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJSON(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannot unmarsha want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannot unmarsha got %q: %v", got, err)
	}
	if diff := cmp.Diff(jw, jg); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() {
		_ = got.Body.Close()
	})

	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}

	if got.StatusCode != status {
		t.Fatalf("want status %d, got %d", status, got.StatusCode)
	}

	if len(gb) == 0 && len(body) == 0 {
		// 어느 쪽이든 응답 바디가 없으므로 AssertJSON 을 호출팔 필요가 없다.
		return
	}
	AssertJSON(t, body, gb)
}

func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read file %q: %v", path, err)
	}
	return bt
}
