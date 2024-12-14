package models

type LoginRequestBody struct {
	Email    string `gorm:"unique" validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type SignUpRequestBody struct {
	FirstName       string `validate:"required,excludesall= " json:"name"`
	LastName        string `validate:"required,nameOrInitials" json:"last_name"`
	Email           string `gorm:"unique" validate:"required,email" json:"email"`
	Password        string `validate:"required,min=8,password" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirmpassword"`
	Phone           string `json:"phone" validate:"required,numeric,len=10"`
}
type AddressRequestBody struct {
	//UserID     uint   `validate:"required"`
	//User       User   `gorm:"foriegnkey:UserID;references:ID"`
	Country    string `json:"country" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50,alpha"`
	State      string `json:"state" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50,alpha"`
	District   string `json:"district" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50,alpha"`
	StreetName string `json:"street_name" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50,alpha"`
	PinCode    string `json:"pin_code" validate:"required,numeric"`
	Phone      string `json:"phone" validate:"required,numeric,len=10"`
	Default    bool   `json:"Default" `
}
type AdminLoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CategoryRequestBody struct {
	//ID           uint   `gorm:"primary key" json:"id"`
	CategoryName string ` gorm:"unique" json:"category_name" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50"`
	Description  string `json:"category_description" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=100"`
	ImageUrl     string `json:"category_imageUrl" validate:"required,max=100,excludesall= "`
}
type ProductRequestBody struct {
	//ID         uint ` json:"id"`
	//CategoryID   uint   `json:"category_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
	//Category    Category `gorm:"foriegnkey:CategoryID;references:ID"`
	ProductName string  `json:"product_name" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=50"`
	Description string  `json:"product_description" validate:"required,no_leading_trailing_spaces,no_repeating_spaces,max=100"`
	ImageUrl    string  `json:"product_imageUrl" validate:"required,excludesall= "`
	Price       float64 ` json:"price" validate:"required,gt=0"`
	Stock       uint    `json:"stock" validate:"required"`
	Popular     bool    `json:"popular" validate:"required"`
	Size        string  ` json:"size" validate:"required"`
}
type CartRequestBody struct {
	ProductID string `json:"product_id" validate:"required,numeric"`
	Quantiy uint `json:"quantity" validate:"required,numeric"`
}

type OrderRequestBody struct {
	AddressID  string   `json:"address_id" validate:"required"`
	//CouponCode string `json:"coupon_code"`
}

type OrderItemRequestBody struct {
	OrderID  string   `json:"order_id" validate:"required"`
	//CouponCode string `json:"coupon_code"`
}