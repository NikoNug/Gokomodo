package models

type Seller struct {
	SellerID int    `json:"seller_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Pickup   string `json:"pickup"`
}

type Buyer struct {
	BuyerID           int    `json:"buyer_id"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	Password          string `json:"password"`
	Alamat_Pengiriman string `json:"alamat_pengiriman"`
}
