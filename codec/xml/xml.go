package xml

import (
	"encoding/xml"
)

type Codec struct {
}

func (b *Codec) MIMEType() string {
	return "application/xml"
}

// 将结构体编码为xml的字节数组
func (x *Codec) Marshal(v any) (data []byte, err error) {
	return xml.Marshal(v)

}

// 将xml的字节数组解码为结构体
func (x *Codec) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
