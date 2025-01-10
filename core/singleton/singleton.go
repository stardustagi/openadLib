package singleton

import "sync"

// SingletonInterface 定义了单例接口
type SingletonInterface[T any] interface {
	Get() T
}

// Singleton 实现了 SingletonInterface
type Singleton[T any] struct {
	value T
	once  sync.Once
}

// Get 返回单例的值
func (s *Singleton[T]) Get() T {
	return s.value
}
