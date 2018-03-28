package internal

import (
	"github.com/kihamo/shadow/components/config"
)

type WatcherItem struct {
	config.Watcher

	source  string
	watcher config.Watcher
}

func NewWatcherItem(watcher config.Watcher, source string) *WatcherItem {
	return &WatcherItem{
		watcher: watcher,
		source:  source,
	}
}

func (w *WatcherItem) Source() string {
	return w.source
}

func (w *WatcherItem) Keys() []string {
	return w.watcher.Keys()
}

func (w *WatcherItem) Callback(key string, new interface{}, old interface{}) {
	w.watcher.Callback(key, new, old)
}
