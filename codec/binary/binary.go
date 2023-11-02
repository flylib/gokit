package binary

import (
	"github.com/davyxu/goobjfmt"
)

type Codec struct {
}

func (b *Codec) MIMEType() string {
	return "application/binary"
}

func (b *Codec) Marshal(v any) (data []byte, err error) {
	return goobjfmt.BinaryWrite(v)
}

func (b *Codec) Unmarshal(data []byte, msgObj any) error {
	return goobjfmt.BinaryRead(data, msgObj)
}
