package constant

import "time"

const (
	RedisAccountStatus = "account:status:%s"
	RedisDefaultExp    = 24 * 7 * time.Hour
)
