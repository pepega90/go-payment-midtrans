package models

import "github.com/midtrans/midtrans-go"

type Data struct {
	Tipe  string                 `json:"tipe,omitempty"`
	Items []midtrans.ItemDetails `json:"items,omitempty"`
}

func (d *Data) GetTotal() int64 {
	var total int64
	for _, v := range d.Items {
		total += v.Price * int64(v.Qty)
	}
	return total
}
