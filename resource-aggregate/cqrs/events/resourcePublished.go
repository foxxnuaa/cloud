package events

import (
	"github.com/go-ocf/cloud/resource-aggregate/pb"
	"github.com/go-ocf/kit/net/http"
)

type ResourcePublished struct {
	pb.ResourcePublished
}

func (e ResourcePublished) Version() uint64 {
	return e.EventMetadata.Version
}

func (e ResourcePublished) Marshal() ([]byte, error) {
	return e.ResourcePublished.Marshal()
}

func (e *ResourcePublished) Unmarshal(b []byte) error {
	return e.ResourcePublished.Unmarshal(b)
}

func (e ResourcePublished) EventType() string {
	return http.ProtobufContentType(&pb.ResourcePublished{})
}

func (e ResourcePublished) AggregateId() string {
	return e.Id
}
