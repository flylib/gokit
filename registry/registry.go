package registry

import "time"

type IRegistry interface {
	Connect(addr ...string) error
	Close() error
	Register(App) error
	Deregister(App) error
	GetAllByGroup(group string) ([]App, error)
	GetAllByKey(key string) ([]App, error)
	Watch(func(action Action, metadata App)) error
}

type App interface {
	Group() string
	Version() string
	Timestamp() time.Duration
	Key() string  //key
	Body() []byte //value
	UnmarshalTo(v any) error
}

type Action int8

const (
	ActionAdd Action = iota + 1
	ActionDel
	ActionUpdate
)
