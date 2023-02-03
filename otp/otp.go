package otp

import (
	"fmt"
	"time"
	"math/rand"
	"strings"
	"os"

	"gorm.io/gorm"
)


type OtpModel struct{
	gorm.Model  //This auto generates CreatedAt and ID fields
	UserEmail	string
	Token	string	
}

func validateEmail(email string)(bool){
	err := strings.Contains(email, "@")
	if !err{
		fmt.Println("add a valid email address.")	
		os.Exit(1)
		
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
		os.Exit(1)
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

func ValidateOTP(db *gorm.DB, userEmail string, token string)(err error){
	var foundOTP OtpModel

	validateEmail(userEmail)


	err = db.Model(OtpModel{}).Where("user_email = ?", userEmail).Where("token = ?", token).Find(&foundOTP).Error
	if err !=  nil{
		return 
	}
	fmt.Printf("%s One Time Password Validated ", userEmail)
	return 
}