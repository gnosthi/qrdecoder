package main

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestDecodeQR tests QRDecoding functions.
func TestDecodeQR(t *testing.T) {
	_, dErr := DecodeQR("")
	assert.Errorf(t, dErr, "no such file")

	q, qErr := DecodeQR("test_qr.png")
	if qErr != nil {
		t.Errorf("%v",qErr)
	}

	assert.Equal(t, "test_qr.png", q.Location)

	for _, data := range q.DecodedData {
		assert.Equal(t, "This is a test of the emergency broadcast system.", string(data.Payload))
	}
}

func TestMain(m *testing.M) {
	os.Args = []string{"test_qr.png"}
	os.Exit(m.Run())
}
