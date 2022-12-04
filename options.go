package rabbitmq

import "time"

type Options struct {
	retryTimes     int64
	backOffSeconds time.Duration
}

type Option func(*Options)

func WithRetryTimes(retryTimes int64) Option {
	return func(o *Options) {
		o.retryTimes = retryTimes
	}
}

func WithBackOffSeconds(backOffSeconds time.Duration) Option {
	return func(o *Options) {
		o.backOffSeconds = 2
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		retryTimes:     retryTimes,
		backOffSeconds: backOffSeconds,
	}

	for _, o := range opts {
		o(options)
	}

	return options
}
