package design_mode


// 抽离共同属性部分
type People struct {
	name string
	age  int
}


func (p *People) Name() string {
	return p.name
}


func (p *People) Age() int {
	return p.age
}


// 继承公共属性部分
type Coder struct {
	People
	skill string
}


func (c Coder) Skill() string {
	return c.skill
}


func NewCoder(name, skill string, age int) *Coder {
	return &Coder{
		People: People{
			name: name,
			age:  age,
		},
		skill: skill,
	}
}
