package watcher

import (
	"golang.org/x/sys/unix"
	"time"
)

const (
	KindCreate  = "CREATE"
	KindDelete  = "DELETE"
	KindChange  = "CHANGE"
	KindUnknown = "UNKNOWN"
)

type Event struct {
	Path         string
	Mask         uint64
	Created      time.Time
	LastModified time.Time
}

func (e Event) Kind() string {
	masks := map[uint64]string{
		unix.FAN_CREATE:     KindCreate,
		unix.FAN_MODIFY:     KindChange,
		unix.FAN_ATTRIB:     KindChange,
		unix.FAN_DELETE:     KindDelete,
		unix.FAN_MOVED_TO:   KindCreate,
		unix.FAN_MOVED_FROM: KindDelete,
	}

	for m, desc := range masks {
		if e.Mask&m != 0 {
			return desc
		}
	}

	return KindUnknown
}

func debounceEvent(old, new Event) Event {
	switch new.Kind() {
	case KindCreate:
		if old.Kind() == KindDelete {
			// A previously deleted file was recreated. Therefore, the event must be rewritten to a change type
			old.Mask = unix.FAN_MODIFY
		} else {
			old.Mask = new.Mask
		}
		old.LastModified = new.Created
	case KindDelete:
		old.Mask = new.Mask
		old.LastModified = new.Created
	case KindChange:
		if old.Kind() == KindDelete {
			// Sometimes on creation of a file a "CHANGE" event gets emitted instead of a "CREATE".
			// We handle it like in the "CREATE" case
			old.Mask = unix.FAN_MODIFY
		}
		old.LastModified = new.Created
	case KindUnknown:
		old.LastModified = new.Created
	}

	return old
}
