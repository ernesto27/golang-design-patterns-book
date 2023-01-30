package chainresponsability

import (
	"strings"
	"testing"
)

func TestCreateDefaultChain(t *testing.T) {
	w := myTestWriter{}

	wl := WriteLogger{Writer: &w}
	second := SecondLogger{NextChain: &wl}
	chain := FirstLogger{NextChain: &second}

	t.Run("3 loggers, 2 of them writes to console, second only if it founds "+"the word 'hello', third writes to some variable if second found 'hello'",
		func(t *testing.T) {
			chain.Next("message that breaks the chain\n")
			if w.receivedMessage != nil {
				t.Fatal("Last link should not receive any message")
			}
			chain.Next("Hello\n")
			if !strings.Contains(*w.receivedMessage, "Hello") {
				t.Fatal("Last link didn't received expected message")
			}
		})
}
