package design_mode

import "errors"

type Kind int


const (
	Normal Kind = iota
	Mock
)


type Loader interface {
	Load() string
}


// 实现两个struct， 继承接口Load
type NormalLoader struct {
	Content string
}


type MockLoader struct {
	Content string
}


func (loader *NormalLoader) Load() string {
	return loader.Content
}


func (loader *MockLoader) Load() string {
	return "Mock " + loader.Content
}


// GenerateLoader 简单工厂模式       使一个类的实例化延迟到其子类, 定义一个用于创建对象的接口, 让子类决定将哪一个类实例化
func GenerateLoader(k Kind, content string) (Loader, error) {
	switch k {
	case Normal:
		return &NormalLoader{content}, nil
	case Mock:
		return &MockLoader{content}, nil
	default:
		return nil, errors.New("Loader do not supports this ")
	}
}