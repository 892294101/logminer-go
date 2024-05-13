package connector

import "time"

type ConnectConfig struct {
	cacheFileSize    uint64
	cacheGroup       int
	flushPeriodTime  time.Duration
	flushMultipleOps int
}

type ConnectOption func(*ConnectConfig) error

func NewConnectorConfig() *ConnectConfig {
	return new(ConnectConfig)
}

func SetCacheSize(cacheSize uint64) ConnectOption {
	return func(c *ConnectConfig) error {

		return nil
	}
}
