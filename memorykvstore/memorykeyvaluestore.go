package memorykvstore

import (
	"bytes"
	"github.com/SomethingBot/discordingestor/ingestor"
	"io"
	"sync"
)

//MemoryKeyValueStore is a KV store that stores []byte(s)
type MemoryKeyValueStore struct {
	store     map[string][]byte
	storeLock sync.RWMutex
}

//New *MemoryKeyValueStore is created and returned
func New() *MemoryKeyValueStore {
	return &MemoryKeyValueStore{
		store: make(map[string][]byte),
	}
}

//Open MemoryKeyValueStore, this is a no-op
func (kvs *MemoryKeyValueStore) Open() error {
	return nil
}

//Close MemoryKeyValueStore, this is a no-op
func (kvs *MemoryKeyValueStore) Close() error {
	return nil
}

//Set at Key, the contents Read from io.Reader; contents not set if errored
func (kvs *MemoryKeyValueStore) Set(key string, reader io.Reader) error {
	kvs.storeLock.Lock()
	defer kvs.storeLock.Unlock()
	var err error
	kvs.store[key], err = io.ReadAll(reader)
	return err
}

//Get at Key and copy to io.Reader
func (kvs *MemoryKeyValueStore) Get(key string) (io.Reader, error) {
	kvs.storeLock.RLock()
	defer kvs.storeLock.RUnlock()
	val, ok := kvs.store[key]
	if !ok {
		return nil, ingestor.ErrorKeyDoesNotExist
	}
	return bytes.NewReader(val), nil
}
