package models

type Order struct {
	OrderID                    int    `json:"order_id"`
	BuyerID                    int    `json:"buyer_id"`
	SellerID                   int    `json:"seller_id"`
	DeliverySourceAddress      string `json:"delivery_source_address"`
	DeliveryDestinationAddress string `json:"delivery_destination_address"`
	Status                     bool   `json:"status"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}
