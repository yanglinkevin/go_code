package design_mode

import "sync"

type Instance struct {
	Name string
}


var (
	instance *Instance
	once     sync.Once
)


// GetInstanceWithSyncOnceInit 双重检查机制
func GetInstanceWithSyncOnceInit() *Instance {
	once.Do(func() {
		// Lock and Init
		instance = new(Instance)
		instance.Name = "First Init With Lock"
	})
	return instance
}



/* 饿汉 有性能问题
type Instance struct {
	Name string
}


var instance *Instance


func init() {
	instance = new(Instance)
	instance.Name = "Hunger Init"
}


// GetInstanceWithHungerInit 饿汉式初始化
func GetInstanceWithHungerInit() *Instance {
	return instance
}
 */





// 懒汉模式 有并发安全问题
/*
type Instance struct {
	Name string
}


var instance *Instance


// GetInstanceWithLazyInit 饿汉式初始化
func GetInstanceWithLazyInit() *Instance {
	if instance == nil {
		instance = new(Instance)
	}
	return instance
}
*/
