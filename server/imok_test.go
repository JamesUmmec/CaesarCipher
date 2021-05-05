package local_server

import (
	"testing"
	"time"
)

// The channel will be create outer and be related to communication
// with the frontend site in normal use. Here is just for test.
func TestWatcher(t *testing.T) {
	watcher := ImokWatch(500*time.Millisecond, 5)
	watcher.Watch(make(chan bool))
}
