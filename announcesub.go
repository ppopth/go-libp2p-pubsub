package pubsub

import (
	"context"
)

const (
	AnnounceSubID_v10 = protocol.ID("/announcesub/1.0.0")
)

// Defines the default announcesub parameters.
var (
	AnnounceSubD                 = 6
	AnnounceSubDlo               = 5
	AnnounceSubDhi               = 12
	AnnounceSubHistoryLength     = 5
	AnnounceSubHistoryGossip     = 3
	AnnounceSubDlazy             = 6
	AnnounceSubHeartbeatInterval = 1 * time.Second
	AnnounceSubFanoutTTL         = 60 * time.Second
)

// AnnounceSubParams defines all the announcesub specific parameters.
type AnnounceSubParams struct {
}
