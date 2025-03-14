package products

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	CategoryID  int     `json:"category_id" gorm:"not null"`
	Image       string  `json:"image" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
}
