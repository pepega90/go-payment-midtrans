package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go_payment_midtrans/models"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

// TODO
// BELI LEWAT CSTORE
// BELI LEWAT CREDIT CARD

func (app *Config) BeliLewatBANK(c *gin.Context) {
	var beli models.Data
	m := coreapi.Client{}
	m.New(midtrans.ServerKey, midtrans.Sandbox)

	if err := c.ShouldBindJSON(&beli); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid data",
		})
		return
	}
	rand.Seed(time.Now().UnixNano())
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(rand.Intn(20000-10000) + 10000),
			GrossAmt: int64(beli.GetTotal()),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: "Aji",
			LName: "Mustofa",
			Email: "pepeg2a@gmail.com",
			Phone: "085123124123",
		},
		Items: &beli.Items,
	}

	switch beli.Tipe {
	case "bca":
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank:     midtrans.Bank(beli.Tipe),
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
	case "permata":
		chargeReq.BankTransfer = &coreapi.BankTransferDetails{
			Bank:     midtrans.Bank(beli.Tipe),
			VaNumber: "1234567890",
			Permata: &coreapi.PermataBankTransferDetail{
				RecipientName: chargeReq.CustomerDetails.FName + chargeReq.CustomerDetails.LName,
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

	res, err := m.ChargeTransaction(chargeReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Message)
		return
	}
	c.JSON(http.StatusOK, res)
}
