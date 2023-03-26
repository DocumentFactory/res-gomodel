package etcdutils

import (
	"context"
	"errors"
	"log"
	"strconv"
)

type EtcdHelper struct {
	ctx context.Context
	db  *DB
}

func NewEtcdHelper(ctx context.Context, url string, keyprefix string) (*EtcdHelper, error) {
	c := EtcdHelper{}
	db, err := GetDB(ctx, url, keyprefix)
	if err != nil {
		return nil, err
	}
	c.ctx = ctx
	c.db = db
	return &c, nil
}

func GetDB(ctx context.Context, url string, keyprefix string) (*DB, error) {

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
		return nil, errors.New("key not found")
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

func (eh *EtcdHelper) Increment(key string) error {
	bytes, err := eh.Get(key)
	if err != nil {
		bytes = []byte("0")
		err = eh.Put(key, bytes)
	}
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}

	bytes = []byte(strconv.Itoa(i + 1))
	err = eh.Put(key, bytes)
	if err != nil {
		return err
	}
	return nil
}
func (eh *EtcdHelper) Decrement(key string) error {
	bytes, err := eh.Get(key)
	if err != nil {
		bytes = []byte("0")
		err = eh.Put(key, bytes)
	}
	if err != nil {
		return err
	}

	i, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}
	if i == 0 {
		return nil
	}
	bytes = []byte(strconv.Itoa(i - 1))
	err = eh.Put(key, bytes)
	if err != nil {
		return err
	}
	return nil
}

func (eh *EtcdHelper) Watch(key string, fn func(old, new interface{})) (context.CancelFunc, error) {
	watchCtx, watchCancel := context.WithCancel(context.Background())

	return watchCancel, eh.db.WatchKey(watchCtx, key, fn)

}

func (eh *EtcdHelper) Close() {
	eh.db.Close()
}
