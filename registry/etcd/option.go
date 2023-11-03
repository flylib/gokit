package etcd

import "time"

type Option func(*registry)

func WithAuth(username, password string) Option {
	return func(r *registry) {
		r.Config.Username = username
		r.Config.Password = password
	}
}

func WithDialTimeout(d time.Duration) Option {
	return func(r *registry) {
		r.Config.DialTimeout = d
	}
}

func WithRegistryTTL(d time.Duration) Option {
	return func(r *registry) {
		r.TTL = d
	}
}

func WithRegistryInterval(d time.Duration) Option {
	return func(r *registry) {
		r.Interval = d
	}
}
