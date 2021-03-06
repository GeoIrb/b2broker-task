package mq

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testURL   = "amqp://guest:guest@localhost:5672/"
	testQueue = "queue-test"
)

var (
	testMessage = []byte("test-message")
)

func TestConnect(t *testing.T) {
	mq, err := NewRabbitMQ(testURL)
	assert.NoError(t, err, "connect to mq")
	assert.NotNil(t, mq, "connect to mq")
	assert.NotNil(t, mq.connection, "connect to mq: connection")
	assert.NotNil(t, mq.channel, "connect to mq: channel")
	mq.Shoutdown()
}

func TestQueue(t *testing.T) {
	mq, err := NewRabbitMQ(testURL)
	assert.NoError(t, err, "connect to mq")
	defer mq.Shoutdown()

	send, err := mq.Publisher(testQueue)
	assert.NoError(t, err, "publisher")

	err = send(context.Background(), testMessage)
	assert.NoError(t, err, "send message")

	err = mq.Consumer(testQueue, testHandler(t))
	assert.NoError(t, err, "add consumer for new queue")

	err = mq.Consumer(testQueue, testHandler(t))
	assert.Error(t, err, "add consumer for existing queue")

	mq.ListenAndServe()
}

func testHandler(t *testing.T) Handler {
	return func(ctx context.Context, data []byte) {
		assert.Equal(t, testMessage, data)
	}
}
