package entities

const tbTransaction = "transaction"

type Transaction struct {
	ID           *int    `gorm:"column:id" json:"id"`
	Address      *string `gorm:"column:address" json:"address"`
	CreatedAt    *string `gorm:"column:created_at" json:"created_at"`
	ProductId    *int    `gorm:"column:product_id" json:"product_id"`
	ProductName  *string `gorm:"column:productName" json:"productName"`
	ProductImage *string `gorm:"column:productImage" json:"productImage"`
	UserId       *int    `gorm:"column:user_id" json:"user_id"`
	UserName     *string `gorm:"column:userName" json:"userName"`
	Amount       *int    `gorm:"column:amount" json:"amount"`
}

type TransactionInsert struct {
	ID        *int    `gorm:"column:id" json:"id"`
	Address   *string `gorm:"column:address" json:"address"`
	ProductId *string `gorm:"column:product_id" json:"product_id"`
	UserId    *int    `gorm:"column:user_id" json:"user_id"`
	Amount    *int    `gorm:"column:amount" json:"amount"`
}

func (*Transaction) TableName() string {
	return tbTransaction
}

type TransactionBatch struct {
	Data *[]TransactionInsert
}

func (*TransactionInsert) TableName() string {
	return tbTransaction
}
