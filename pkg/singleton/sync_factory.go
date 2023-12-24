package singleton

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type SecureStore interface {
	Value(key string) string
}

// Abstract method to return value
func GetSecretValue(store SecureStore, key string) string {
	return store.Value(key)
}

// Hide the type from usage outside the package
type secretStore struct {
	kv map[string]string
}

func (s *secretStore) Value(k string) string {
	return s.kv[k]
}

var once sync.Once
var instance SecureStore

func NewSecret() SecureStore {
	once.Do(func() {
		fmt.Println("Loading secrets...")
		kv, e := loadSecrets(".\\secrets.txt")
		ss := secretStore{kv}
		if e == nil {
			ss.kv = kv
		}
		instance = &ss
	})
	return instance
}

func loadSecrets(path string) (map[string]string, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]string{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v := scanner.Text()
		result[k] = v
	}

	return result, nil
}
