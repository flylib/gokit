package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type registry struct {
	client   *clientv3.Client
	cfg      clientv3.Config
	lease    clientv3.Lease
	TTL      time.Duration
	Interval time.Duration
	leaseID  clientv3.LeaseID
	story    clientv3.KV
}

func NewRegistry(addrs []string, options ...Option) {
	rgs := registry{
		cfg: clientv3.Config{
			Endpoints: addrs,
		},
		TTL:      time.Second * 25,
		Interval: time.Second * 20,
	}
	for _, f := range options {
		f(&rgs)
	}
	return &rgs
}

func (r *registry) Connect(addr ...string) error {
	r.cfg.Endpoints = addr
	cli, err := clientv3.New(r.cfg)
	if err != nil {
		return err
	}
	r.client = cli

	//创建一个租约
	r.lease = clientv3.NewLease(r.client)

	//设置租约时间
	leaseResp, err := r.lease.Grant(context.TODO(), int64(r.TTL/time.Second))
	if err != nil {
		return err
	}

	//设置续租
	leaseRespChan, err := r.lease.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return err
	}
	go func() {
		for info := range leaseRespChan {
			fmt.Println("lease success!!!", info.TTL, " id ", info.ID)
		}
	}()

	r.leaseID = leaseResp.ID

	r.story = clientv3.NewKV(r.client)
	return err
}

func (r *registry) Close() error {
	r.lease.Revoke(context.Background(), r.leaseID)
	return r.client.Close()
}

func (r *registry) Register(kv KV) error {
	_, err := r.story.Put(context.TODO(), key, val, clientv3.WithLease(r.leaseID))
	return err
}

func (r *registry) Deregister(kv KV) error {
	_, err := r.story.Delete(context.TODO(), key, val, clientv3.WithLease(r.leaseID))
	return err
}

func (r *registry) GetAllByGroup(group string) ([]KV, error) {
	r.story.Get()
}

func (r *registry) GetAllByKey(key string) ([]KV, error) {
	//TODO implement me
	panic("implement me")
}

func (r *registry) Watch(f func(action Action, metadata KV)) error {
	//TODO implement me
	panic("implement me")
}
