package design_pattern

// 顶层操作接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 顶层创建操作的接口
type OperatorFactory interface {
	CreateFactory() Operator
}

// 顶层数据存放
type BaseOperator struct {
	a, b int
}

func (bo *BaseOperator) SetA(a int) {
	bo.a = a
}

func (bo *BaseOperator) SetB(b int) {
	bo.b = b
}

// 操作实类 1
type SubOperator struct {
	*BaseOperator
}

func (so *SubOperator) Result() int {
	return so.b - so.a
}

type SubOperatorFactory struct {
}

func (SubOperatorFactory) CreateFactory() Operator {
	return &SubOperator{&BaseOperator{}}
}

// 操作实类 2
type AddOperator struct {
	*BaseOperator
}

func (ao *AddOperator) Result() int {
	return ao.b + ao.a
}

type AddOperatorFactory struct {
}

func (AddOperatorFactory) CreateFactory() Operator {
	return &AddOperator{&BaseOperator{}}
}
