package modelq

type Product struct {
	// Name of the product
	Name string `json:"name"`

	// Code of the product
	Code string `json:"code"`

	// Word category of the product
	Category string `json:"category"`
}

type ProductImport struct {
	Products []Product `json:"products"`
}