package storage

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type SetGetFlusher interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
	Flush()
	Items() map[string]cache.Item
}
