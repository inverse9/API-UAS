package entities

const tbTransaction = "transaction"

type Transaction struct {
	ID          *int    `gorm:"column:id" json:"id"`
	Address     *string `gorm:"column:address" json:"address"`
	CreatedAt   *string `gorm:"column:created_at" json:"created_at"`
	ProductId   *int    `gorm:"column:product_id" json:"product_id"`
	ProductName *string `gorm:"column:productName" json:"productName"`
	UserId      *int    `gorm:"column:user_id" json:"user_id"`
	UserName    *string `gorm:"column:userName" json:"userName"`
}

type TransactionInsert struct {
	ID        *int    `gorm:"column:id" json:"id"`
	Address   *string `gorm:"column:address" json:"address"`
	ProductId *int    `gorm:"column:product_id" json:"product_id"`
	UserId    *int    `gorm:"column:user_id" json:"user_id"`
}

func (*Transaction) TableName() string {
	return tbTransaction
}

func (*TransactionInsert) TableName() string {
	return tbTransaction
}
