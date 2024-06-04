package entities

const tbUser = "user"

type User struct {
	ID        *int    `gorm:"column:id" json:"id"`
	Name      *string `gorm:"column:name" json:"name"`
	CreatedAt *string `gorm:"column:created_at" json:"created_at"`
	ShopId    *int    `gorm:"column:shopId" json:"shopId"`
}

type UserInsert struct {
	ID       *int    `gorm:"column:id" json:"id"`
	Name     *string `gorm:"column:name" json:"name"`
	Password *string `gorm:"column:password" json:"password"`
}

func (*User) TableName() string {
	return tbUser
}

func (*UserInsert) TableName() string {
	return tbUser
}