package singleton

import (
	"testing"
)

type MySingleton struct {
	*SingletonManager
}

func NewMySingleton() *MySingleton {
	return &MySingleton{
		SingletonManager: NewSingletonManager(),
	}
}

// CreateOrGetIntSingleton 创建或获取 int 类型的单例
func (ms *MySingleton) CreateOrGetIntSingleton(key string, initFunc func() int) SingletonInterface[int] {
	instance := ms.GetOrCreateInstance(key, func() interface{} {
		return &Singleton[int]{value: initFunc()}
	})
	return instance.(SingletonInterface[int])
}

func TestSingleton(t *testing.T) {
	mySingletonManager := NewMySingleton()

	// 获取或创建一个 int 类型的单例
	intSingleton := mySingletonManager.CreateOrGetIntSingleton("myInt", func() int {
		return 42
	})

	// 打印单例值
	println(intSingleton.Get())
}
