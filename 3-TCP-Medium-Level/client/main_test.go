package main

import "testing"

func TestMessageParse(t *testing.T) {
	message := "hello"

	// MessageTypeText türünde "hello" mesajı oluştur
	data := createMessage(MessageTypeText, message)

	// Oluşturulan mesajı parse et (oku)
	mtype, mlen, msg := readMessage(data)

	// Mesaj türünün doğruluğunu kontrol et
	if mtype != MessageTypeText {
		t.Errorf("Beklenen mesaj türü %d, alınan %d", MessageTypeText, mtype)
	}

	// Mesaj uzunluğunun doğruluğunu kontrol et
	if mlen != uint32(len(message)) {
		t.Errorf("Beklenen mesaj uzunluğu %d, alınan %d", len(message), mlen)
	}

	// Mesaj içeriğinin doğruluğunu kontrol et
	if string(msg) != message {
		t.Errorf("Beklenen mesaj içeriği %s, alınan %s", message, msg)
	}
}
