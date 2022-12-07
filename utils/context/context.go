package context

import (
	"context"
	"time"
)

func ContextWithTimeout(t time.Duration) context.Context {
	c, _ := context.WithTimeout(context.Background(), t)
	return c
}
