package entities

const tbCart = "cart"

type Cart struct {
	ID          *int    `gorm:"column:id" json:"id"`
	ProductId   *string `gorm:"column:product_id" json:"product_id"`
	ProductName *string `gorm:"column:productName" json:"productName"`
	UserId      *int    `gorm:"column:user_id" json:"user_id"`
	UserName    *string `gorm:"column:userName" json:"userName"`
	Amount      *int    `gorm:"column:amount" json:"amount"`
}

type CartInsert struct {
	ID        *int `gorm:"column:id" json:"id"`
	UserId    *int `gorm:"column:user_id" json:"user_id"`
	ProductId *int `gorm:"column:product_id" json:"product_id"`
	Amount    *int `gorm:"column:amount" json:"amount"`
}

func (*Cart) TableName() string {
	return tbCart
}

func (*CartInsert) TableName() string {
	return tbCart
}
