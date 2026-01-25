package models

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Categories = []Category{
	{
		ID:          1,
		Name:        "Electronics",
		Description: "Devices and gadgets for everyday use",
	},
	{
		ID:          2,
		Name:        "Books",
		Description: "A wide range of books from various genres",
	},
	{
		ID:          3,
		Name:        "Clothing",
		Description: "Apparel and fashion accessories for everyone",
	},
	{
		ID:          4,
		Name:        "Home & Kitchen",
		Description: "Essentials and decor for home and kitchen",
	},
	{
		ID:          5,
		Name:        "Toys",
		Description: "Fun and educational toys for children of all ages",
	},
}
