package pattern

import "sync"

type singleton struct {
	name string
}

var instance *singleton
var once sync.Once
var mu sync.Mutex

//借助sync.Once实现单例
func GetInstanceByOnce() *singleton {
	once.Do(func() {
		instance = &singleton{name: "ByOnce"}
	})
	return instance
}

//线程安全的懒汉式
//大部分懒汉式实现不加锁。。可能是因为加锁过于重，但不加锁可能会有线程安全问题
func GetInstanceByLazy() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{name: "ByLazy"}
	}
	return instance
}

//双重检测
//减少了不必要的加锁开销
func GetInstanceByDouble() *singleton {
	if instance ==nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{name: "ByLazy"}
		}
		return instance
	}
	return instance
}

//init函数将在包初始化时执行(执行一次)，实例化单例,
func init() {
	instance = &singleton{name: "ByHungry"}
}
func GetInstanceByHungry() *singleton {
	return instance
}
