package etcdutils

import (
	"context"
	"errors"
	"log"
)

var keyprefix string = "/store/"

type EtcdHelper struct {
	ctx context.Context
	db  *DB
}

func NewEtcdHelper(ctx context.Context, url string) (*EtcdHelper, error) {
	c := EtcdHelper{}
	db, err := GetDB(ctx, url)
	if err != nil {
		return nil, err
	}
	c.ctx = ctx
	c.db = db
	return &c, nil
}

func GetDB(ctx context.Context, url string) (*DB, error) {

	return New(ctx, url, Options{
		Logf: func(format string, args ...interface{}) {
			log.Printf(format, args...)
		},
		DeleteAllOnStart: false,
		WatchFunc: func(k []KV) {
			log.Printf("watch: %v", k)
		},
		KeyPrefix: keyprefix})
}

func (eh *EtcdHelper) Get(key string) ([]byte, error) {
	var bytes []byte
	tx := eh.db.Tx(context.Background())
	found, err := tx.Get(key, &bytes)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, errors.New("Key not found")
	}
	return bytes, nil
}

func (eh *EtcdHelper) Put(key string, value []byte) error {
	tx := eh.db.Tx(context.Background())
	err := tx.Put(key, value)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (eh *EtcdHelper) Watch(key string, fn func(old, new interface{})) (error, context.CancelFunc) {
	watchCtx, watchCancel := context.WithCancel(context.Background())

	return eh.db.WatchKey(watchCtx, key, fn), watchCancel

}

func (eh *EtcdHelper) Close() {
	eh.db.Close()
}
