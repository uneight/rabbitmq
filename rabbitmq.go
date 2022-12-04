package rabbitmq

import (
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

const (
	retryTimes     = 5
	backOffSeconds = 2
)

var ErrCannotConnectRabbitMQ = errors.New("cannot connect to rabbit")

func NewRabbitMQConn(dsn string, logger *zap.Logger, opts *Options) (*amqp.Connection, error) {
	var (
		conn   *amqp.Connection
		counts int64
	)

	for {
		connection, err := amqp.Dial(dsn)
		if err != nil {
			logger.Sugar().Error("RabbitMq at %s not ready...\n", dsn)
			counts++
		} else {
			conn = connection

			break
		}

		if counts > opts.retryTimes {
			logger.Sugar().Error(err)

			return nil, ErrCannotConnectRabbitMQ
		}

		logger.Info("Backing off for 2 seconds...")
		time.Sleep(opts.backOffSeconds * time.Second)

		continue
	}

	logger.Info("Connected to RabbitMQ!")

	return conn, nil
}
