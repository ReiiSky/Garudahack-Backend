package models

type AggregateProps struct {
	ProductName string `json:"productName"`
	NamaUMKM    string `json:"sboName"`
	Desc        string `json:"description"`
	Stock       int    `json:"stok"`
	HargaSatuan int    `json:"unitPrice"`
	ID          string `json:"id"`
}
