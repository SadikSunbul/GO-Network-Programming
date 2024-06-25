package main

import "testing"

func TestMessageParse(t *testing.T) {
	message := "hello"

	data := createMessage(MessageTypeText, message)
	mtype, mlen, msg := readMessage(data)
	if mtype != MessageTypeText {
		t.Errorf("expected %d, got %d", MessageTypeText, mtype)
	}
	if mlen != uint32(len(message)) {
		t.Errorf("expected %d, got %d", len(message), mlen)
	}
	if string(msg) != message {
		t.Errorf("expected %s, got %s", message, msg)
	}
}

// go test
