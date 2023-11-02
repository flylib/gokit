package protobuf

import (
	"github.com/gogo/protobuf/proto"
)

type Codec struct {
}

func (b *Codec) MIMEType() string {
	return "application/x-protobuf"
}

func (g *Codec) Marshal(v any) (data []byte, err error) {
	return proto.Marshal(v.(proto.Message))
}

func (g *Codec) Unmarshal(data []byte, v any) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
