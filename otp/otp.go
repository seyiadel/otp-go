package otp

import (
	"fmt"
	"time"
	"math/rand"
	"strings"
	"gorm.io/gorm"
)


type OtpModel struct{
	gorm.Model  //This auto generates CreatedAt and ID fields
	UserEmail	string
	Token	string	
}

func validateEmail(email string)bool{
	err := strings.Contains(email, "@")
	if !err{
		panic("add a valid email address.")
	 }
	return err
}

func Generate(db *gorm.DB, userEmail string, otpLength int)(token string, err error){
	db.AutoMigrate(&OtpModel{})

	validateEmail(userEmail)
	x := generateOTP(otpLength)
	token = x

	err = db.Create(&OtpModel{
			UserEmail: userEmail,
			Token: token,
		}).Error
	
	if err != nil{
		return
	}

	return
}

func generateOTP(otp_length int)(string){
	otp := ""
 
	min := 0
	max := 9

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < otp_length; i++ {
		generatedOtp := fmt.Sprintf("%d",rand.Intn(max - min + 1) + min)
		otp += generatedOtp
	}
	return otp
}
