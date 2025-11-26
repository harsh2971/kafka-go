package orders

type Order struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
