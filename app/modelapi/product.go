package modelapi

// Product is just a product object, used in responses.
type Product struct {
	// Number ID of the product
	ID uint `json:"id"`

	// Name of the product
	Name string `json:"name"`

	// Code of the product
	Code string `json:"code"`

	// Word category of the product
	Category string `json:"category"`
}

// ProductReq represents product without id.
type ProductReq struct {
	// Name of the product
	Name string `json:"name"`

	// Code of the product
	Code string `json:"code"`

	// Word category of the product
	Category string `json:"category"`
}
