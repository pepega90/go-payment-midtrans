package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	token := SnapUIPayment(c)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Nyoba Snap UI",
		"token": token.Token,
	})
}
