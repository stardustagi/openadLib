package singleton

import "sync"

// SingletonManager 管理所有单例实例
type SingletonManager struct {
	instances map[string]interface{}
	mu        sync.Mutex
}

// NewSingletonManager 创建一个新的 SingletonManager 实例
func NewSingletonManager() *SingletonManager {
	return &SingletonManager{
		instances: make(map[string]interface{}),
	}
}

// GetOrCreateInstance 获取或创建一个单例实例
func (sm *SingletonManager) GetOrCreateInstance(key string, initFunc func() interface{}) interface{} {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if instance, exists := sm.instances[key]; exists {
		return instance
	}

	instance := initFunc()
	sm.instances[key] = instance
	return instance
}
