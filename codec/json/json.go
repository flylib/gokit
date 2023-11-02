package json

import (
	"github.com/json-iterator/go" //高性能json编码库
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Codec struct {
}

func (b *Codec) MIMEType() string {
	return "application/json"
}

// 将结构体编码为JSON的字节数组
func (j *Codec) Marshal(v any) (data []byte, err error) {
	return json.Marshal(v)

}

// 将JSON的字节数组解码为结构体
func (j *Codec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
