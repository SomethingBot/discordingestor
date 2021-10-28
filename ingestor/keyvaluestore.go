package ingestor

import (
	"fmt"
	"io"
)

var (
	ErrorKeyDoesNotExist = fmt.Errorf("KeyValueStore: key does not exist")
)

//type Transaction interface {
//	Set(key string, reader io.Reader) error
//	Get(key string) (io.Reader, error)
//}

//KeyValueStore is the store that Ingestor uses to store Rate limit/Other data
type KeyValueStore interface {
	Open() error
	Close() error
	Set(key string, reader io.Reader) error
	Get(key string) (io.Reader, error)
	//Start() (Transaction, error)
	//Commit(transaction Transaction) error
	//Rollback(transaction Transaction) error
}
