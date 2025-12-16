package utils

import "time"

func GetCurrentTime() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
