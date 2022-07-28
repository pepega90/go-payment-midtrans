package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentType interface {
	ViaBCA()
	ViaPermata()
}

type BankPaymentConfig struct {
	Charge coreapi.ChargeReq
}

func NewPayment(v coreapi.CoreapiPaymentType, dt midtrans.TransactionDetails, dc midtrans.CustomerDetails, keranjang *[]midtrans.ItemDetails) *BankPaymentConfig {
	return &BankPaymentConfig{
		Charge: coreapi.ChargeReq{
			PaymentType:        v,
			TransactionDetails: dt,
			CustomerDetails:    &dc,
			Items:              keranjang,
		},
	}
}

func (p *BankPaymentConfig) ViaBCA() {
	p.Charge.BankTransfer = &coreapi.BankTransferDetails{
		Bank:     midtrans.BankBca,
		VaNumber: "12345678901",
		FreeText: &coreapi.BCABankTransferDetailFreeText{
			Inquiry: []coreapi.BCABankTransferLangDetail{
				{
					LangID: "text indonesia",
					LangEN: "text inggris",
				},
			},
			Payment: []coreapi.BCABankTransferLangDetail{
				{
					LangID: "Pembayaran produk",
					LangEN: "Product payment",
				},
			},
		},
	}

}

func (p *BankPaymentConfig) ViaPermata() {
	p.Charge.BankTransfer = &coreapi.BankTransferDetails{
		Bank:     midtrans.BankPermata,
		VaNumber: "1234567890",
		Permata: &coreapi.PermataBankTransferDetail{
			RecipientName: p.Charge.CustomerDetails.FName + p.Charge.CustomerDetails.LName,
		},
		FreeText: &coreapi.BCABankTransferDetailFreeText{
			Inquiry: []coreapi.BCABankTransferLangDetail{
				{
					LangID: "text indonesia",
					LangEN: "text inggris",
				},
			},
			Payment: []coreapi.BCABankTransferLangDetail{
				{
					LangID: "Pembayaran produk",
					LangEN: "Product payment",
				},
			},
		},
	}

}
