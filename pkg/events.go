package pkg

import (
	"context"

	"github.com/threefoldtech/substrate-client"
)

//go:generate mkdir -p stubs

//go:generate zbusc -module node -version 0.0.1 -name events -package stubs github.com/threefoldtech/zos/pkg+Events stubs/events_stub.go

// EventKind describes event kind
type EventKind int

const (
	// EventSubscribed event is always sent when a new connection to tfchain is done
	// this will let the receiver know that we just reconnected to the network hence
	// possible events loss has occurred. Hence the receiver need to make sure
	// it's in sync with the network
	EventSubscribed EventKind = iota
	// EventReceived mean a new event has been received, and need to be handled
	EventReceived
)

// PublicConfigEvent pubic config event received. The type specify if this is just notification
// of the reconnection, or actual event has been received.
type PublicConfigEvent struct {
	Kind         EventKind
	PublicConfig substrate.OptionPublicConfig
}

// ContractCancelledEvent a contract has been cancelled, The type specify if this is just notification
// of the reconnection, or actual event has been received.
type ContractCancelledEvent struct {
	Kind     EventKind
	Contract uint64
	TwinId   uint32
}

// ContractLockedEvent is raised when a contract is locked/unlocked. On locking the Lock flag will be set to true.
// If Kind is EventSubscribed it means event stream has been reconnected and might be events loss. It's up to the
// handler of this event type to make sure contracts are synched with the grid.
type ContractLockedEvent struct {
	Kind     EventKind
	Contract uint64
	TwinId   uint32
	Lock     bool
}

type Events interface {
	PublicConfigEvent(ctx context.Context) <-chan PublicConfigEvent
	ContractCancelledEvent(ctx context.Context) <-chan ContractCancelledEvent
	ContractLockedEvent(ctx context.Context) <-chan ContractLockedEvent
}
