package design_pattern

import "testing"

// 业务里
func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestUseDAOFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBFactory{}
	getMainAndDetail(factory)
	factory = &XMLFactory{}
	getMainAndDetail(factory)
}
