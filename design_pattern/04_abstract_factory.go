package design_pattern

// 最顶层接口
type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

// 顶层抽象工厂
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// 实现工厂1----------------------------------

type RDBMainDAO struct {
}

func (*RDBMainDAO) SaveOrderMain() {
	println("rdb main -> save")
}

type RDBDetailDAO struct {
}

func (*RDBDetailDAO) SaveOrderDetail() {
	println("rdb detail -> save")
}

type RDBFactory struct {
}

func (*RDBFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// 实现工厂2----------------------------------

type XMLMainDAO struct {
}

func (*XMLMainDAO) SaveOrderMain() {
	println("xml main -> save")
}

type XMLDetailDAO struct {
}

func (*XMLDetailDAO) SaveOrderDetail() {
	println("xml detail -> save")
}

type XMLFactory struct {
}

func (*XMLFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
