package quill

import (
	"log/slog"

	"github.com/wagoodman/go-partybus"
)

// SetLogger sets the logger object used for all logging calls.
func SetLogger(logger *slog.Logger) {
}

// SetBus sets the event bus for all library bus publish events onto (in-library subscriptions are not allowed).
func SetBus(b *partybus.Bus) {
}
