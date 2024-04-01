package domain

// 1. 创建领域对象（domain object）
type Customer struct {
	Id        string
	Name      string
	City      string
	ZipCode   string
	BirthDate string
	Status    string // Status 字段表示该客户是 active 还是 inactive
}

// 2. 由于我们没有业务复杂性，我们可以继续定义我们的 repository 接口，
// 这是领域边界（domain boundary）上的次要端口（secondary port）
// 所以可以在 domain 包中定义

// 定义一个端口就像定义一个协议；任何遵循这个协议的组件都应该能够连接到它
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
