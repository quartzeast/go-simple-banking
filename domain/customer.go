package domain

import "github.com/quartzeast/go-simple-banking/dto"

// 1. 创建领域对象（domain object）
type Customer struct {
	Id        string `db:"customer_id"`
	Name      string `db:"name"`
	City      string `db:"city"`
	ZipCode   string `db:"zipcode"`
	BirthDate string `db:"date_of_birth"`
	Status    string `db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		Name:      c.Name,
		City:      c.City,
		ZipCode:   c.ZipCode,
		BirthDate: c.BirthDate,
		Status:    c.statusAsText(),
	}
}

// 2. 定义 CustomerRepository 接口，在业务逻辑边界创建端口
// 定义一个接口就相当于定义了一个协议，只要遵循该协议就可以作为底层的 Repository 使用
// 这个接口定义了一个 contract，这个合约制定了一个约定，它要求 Customer 这个业务逻辑需要一个怎样的底层 Repository
// 根据以下的定义，我们知道，Customer 业务逻辑需要的底层 Repository，必须可以 FindAll（查找所有）和 ById（根据 id 查找）
// 任何提供这两个方法实现的 Repository 都可以作为底层的 Repository 使用
type CustomerRepository interface {
	FindAll(string) ([]Customer, error)
	ById(string) (*Customer, error)
}
