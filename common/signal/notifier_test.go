package signal_test

import (
	"testing"

	. "github.com/make-money-fast/v2ray-core-v5/common/signal"
)

func TestNotifierSignal(t *testing.T) {
	n := NewNotifier()

	w := n.Wait()
	n.Signal()

	select {
	case <-w:
	default:
		t.Fail()
	}
}
