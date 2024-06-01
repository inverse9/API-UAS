package entities

const tbProduct = "product"

type Product struct {
	ID          *int     `gorm:"column:id" json:"id"`
	Name        *string  `gorm:"column:name" json:"name"`
	Description *string  `gorm:"column:description" json:"description"`
	Price       *float32 `gorm:"column:price" json:"price"`
	Image       *string  `gorm:"column:image" json:"image"`
	ShopId      *int     `gorm:"column:shop_id" json:"shop_id"`
	ShopName    *string  `gorm:"column:shopName" json:"shopName"`
}

type ProductInsert struct {
	ID          *int     `gorm:"column:id" json:"id"`
	ShopId      *int     `gorm:"column:shop_id" json:"shop_id"`
	Name        *string  `gorm:"column:name" json:"name"`
	Description *string  `gorm:"column:description" json:"description"`
	Price       *float32 `gorm:"column:price" json:"price"`
	Image       *string  `gorm:"column:image" json:"image"`
}

type ProductBatch struct {
	Data *[]ProductInsert
}

func (*Product) TableName() string {
	return tbProduct
}

func (*ProductInsert) TableName() string {
	return tbProduct
}
