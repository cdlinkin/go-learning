/*
	exercise:

Define the interface:

	type Storage interface {
	 Save(key string, value []byte) error
	 Load(key string) ([]byte, error)
	 Delete(key string) error
	}

Implement three variants:
1. MemoryStorage
2. FileStorage (a file per key in the directory)
3. LoggingStorage, which wraps any other Storage and logs calls (the “decorator” pattern)

Write a function Backup(src, dst Storage, keys []string) error that accepts interfaces and is unaware of specific types.
*/
package stage2

import (
	"fmt"
	"os"
	"path/filepath"
)

type Storage interface {
	Save(key string, value []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
}

// MemoryStorage
type MemoryStorage struct {
	memory map[string][]byte
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		memory: make(map[string][]byte),
	}
}

func (m *MemoryStorage) Save(key string, value []byte) error {
	if key == "" {
		return fmt.Errorf("the key cannot be empty")
	}
	if value == nil {
		return fmt.Errorf("the value cannot be empty")
	}
	m.memory[key] = value
	return nil
}

func (m *MemoryStorage) Load(key string) ([]byte, error) {
	if key == "" {
		return nil, fmt.Errorf("the key cannot be empty")
	}

	value, ok := m.memory[key]
	if !ok {
		return nil, fmt.Errorf("the key-value pair is not loaded")
	}

	return value, nil
}

func (m *MemoryStorage) Delete(key string) error {
	if key == "" {
		return fmt.Errorf("the key cannot be empty")
	}

	delete(m.memory, key)
	return nil
}

// FileStorage
type FileStorage struct {
	path string
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{
		path: path,
	}
}

func (f *FileStorage) Save(key string, value []byte) error {
	joinPath := filepath.Join(f.path, key)

	if err := os.WriteFile(joinPath, value, 0644); err != nil {
		return fmt.Errorf("file write error: %w", err)
	}

	fmt.Println("the file was written successfully.")
	return nil
}

func (f *FileStorage) Load(key string) ([]byte, error) {
	joinPath := filepath.Join(f.path, key)

	value, err := os.ReadFile(joinPath)
	if err != nil {
		return nil, fmt.Errorf("file read error: %w", err)
	}

	return value, nil
}

func (f *FileStorage) Delete(key string) error {
	joinPath := filepath.Join(f.path, key)

	if err := os.Remove(joinPath); err != nil {
		return fmt.Errorf("file remove error: %w", err)
	}

	return nil
}

// LoggingStorage
type LoggingStorage struct {
	storage Storage
}

func NewLoggingStorage(storage Storage) *LoggingStorage {
	return &LoggingStorage{
		storage: storage,
	}
}

func (l *LoggingStorage) Save(key string, value []byte) error {
	fmt.Println("Save called")

	if err := l.storage.Save(key, value); err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func (l *LoggingStorage) Load(key string) ([]byte, error) {
	fmt.Println("Load called")

	value, err := l.storage.Load(key)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	return value, nil
}

func (l *LoggingStorage) Delete(key string) error {
	fmt.Println("Delete called")

	if err := l.storage.Delete(key); err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// Backup
func Backup(src, dst Storage, key []string) error {
	for _, v := range key {
		value, err := src.Load(v)
		if err != nil {
			return fmt.Errorf("error: %w", err)
		}
		if err := dst.Save(v, value); err != nil {
			return fmt.Errorf("error: %w", err)
		}
	}
	return nil
}
