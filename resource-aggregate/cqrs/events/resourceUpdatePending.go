package events

import (
	"github.com/go-ocf/cloud/resource-aggregate/pb"
	"github.com/go-ocf/kit/net/http"
)

type ResourceUpdatePending struct {
	pb.ResourceUpdatePending
}

func (e ResourceUpdatePending) Version() uint64 {
	return e.EventMetadata.Version
}

func (e ResourceUpdatePending) Marshal() ([]byte, error) {
	return e.ResourceUpdatePending.Marshal()
}

func (e *ResourceUpdatePending) Unmarshal(b []byte) error {
	return e.ResourceUpdatePending.Unmarshal(b)
}

func (e ResourceUpdatePending) EventType() string {
	return http.ProtobufContentType(&pb.ResourceUpdatePending{})
}

func (e ResourceUpdatePending) AggregateId() string {
	return e.Id
}
