package entities

const tbShop = "shop"

type Shop struct {
	ID       *int    `gorm:"column:id" json:"id"`
	Name     *string `gorm:"column:name" json:"name"`
	Address  *string `gorm:"column:address" json:"address"`
	Picture  *string `gorm:"column:picture" json:"picture"`
	UserId   *int    `gorm:"column:user_id" json:"user_id"`
	UserName *string `gorm:"column:userName" json:"userName"`
}

type ShopInsert struct {
	ID      *int    `gorm:"column:id" json:"id"`
	Name    *string `gorm:"column:name" json:"name"`
	Address *string `gorm:"column:address" json:"address"`
	Picture *string `gorm:"column:picture" json:"picture"`
	UserId  *int    `gorm:"column:user_id" json:"user_id"`
}

func (*Shop) TableName() string {
	return tbShop
}

func (*ShopInsert) TableName() string {
	return tbShop
}
