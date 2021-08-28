package design_mode

// 基础接口
type BaseItem interface {
	Name() string
}


// 基础结构体
type Item struct {
	name string
}


func (i *Item) Name() string {
	return i.name
}


// 装饰结构体
type ItemDecorator struct {
	BaseItem
	Type string
}


func (id *ItemDecorator) Name() string {
	return id.BaseItem.Name() + ": " + id.Type
}


func CreateItemDecorator(b BaseItem, t string) *ItemDecorator {
	return &ItemDecorator{b, t}
}