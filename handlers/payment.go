package handlers

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go_payment_midtrans/config"
	"github.com/go_payment_midtrans/models"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

// TODO
// BELI LEWAT CSTORE
// BELI LEWAT CREDIT CARD

func SnapUIPayment(c *gin.Context) *snap.Response {
	s := snap.Client{}
	s.New(midtrans.ServerKey, midtrans.Sandbox)

	custAddress := &midtrans.CustomerAddress{
		FName:       "Sigit",
		LName:       "Ardianto",
		Phone:       "081234567890",
		Address:     "VTE B69 No.69",
		City:        "Tangerang",
		Postcode:    "15560",
		CountryCode: "IDN",
	}

	req := &snap.Request{TransactionDetails: midtrans.TransactionDetails{
		OrderID:  "MID-GO-ID-" + example.Random(),
		GrossAmt: 200000,
	},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    "Aji",
			LName:    "Mustofa",
			Email:    "aji@handsome.com",
			Phone:    "081234567890",
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "Dildo",
				Price: 200000,
				Qty:   1,
				Name:  "Someitem",
			},
		},
	}

	resp, err := s.CreateTransaction(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.GetMessage(),
		})

	}
	// c.JSON(http.StatusOK, resp)
	return resp
}

func BeliLewatBANK(c *gin.Context) {
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

	transaksi := config.NewPayment(
		coreapi.PaymentTypeBankTransfer,
		midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(rand.Intn(20000-10000) + 10000),
			GrossAmt: int64(beli.GetTotal()),
		},
		midtrans.CustomerDetails{
			FName: "Aji",
			LName: "Mustofa",
			Email: "pepeg2a@gmail.com",
			Phone: "085123124123",
		},
		&beli.Items,
	)

	switch beli.Tipe {
	case "bca":
		transaksi.ViaBCA()
	case "permata":
		transaksi.ViaPermata()
	}

	res, err := m.ChargeTransaction(&transaksi.Charge)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Message)
		return
	}
	c.JSON(http.StatusOK, res)
}
