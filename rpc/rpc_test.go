package rpc_test

import (
	"css-var-lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Method bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Method: true})
	if expected != actual {
		t.Fatalf("Expected: %s ; got: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMsg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMsg))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if contentLength != 15 {
		t.Fatalf("wrong content length returned expected: %d; got= %d", 15, contentLength)
	}
	if method != "hi" {
		t.Fatalf("wrong method: got %s, expected %s", method, "hi")
	}
}
