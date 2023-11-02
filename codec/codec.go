package codec

type MIMEType string

const (
	MIMETypeBinary   MIMEType = "application/binary"
	MIMETypeXml      MIMEType = "application/xml"
	MIMETypeJson     MIMEType = "application/json"
	MIMETypeProtobuf MIMEType = "application/x-protobuf"
)

type ICodec interface {
	MIMEType() string
	Marshal(v any) (data []byte, err error)
	Unmarshal(data []byte, v any) error
}
