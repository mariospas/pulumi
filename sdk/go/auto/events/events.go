package events

import "github.com/mariospas/pulumi/sdk/v3/go/common/apitype"

type EngineEvent struct {
	apitype.EngineEvent
	Error error
}
