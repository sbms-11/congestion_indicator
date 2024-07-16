package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SdampleFanc(c *gin.Context) {
	log.Println("Hello, im Gopher")
	c.JSON(http.StatusOK, gin.H{"message": "Hi!"})
}

func GetHomeInfo(c *gin.Context) {

}

func GetUserInfo(c *gin.Context) {

}

func GetShopStatus(c *gin.Context) {

}

func MakeNewReservation(c *gin.Context) {

}
