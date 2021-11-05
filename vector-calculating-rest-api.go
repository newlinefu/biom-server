package main

import (
	"github.com/gin-gonic/gin"
)

func GetCalculatedVector(c *gin.Context) {
	var registrationObjects RegistrationObject
	err := c.BindJSON(&registrationObjects)
	if err != nil {
		c.Error(err)
	}
	information, vector := CalculateVectorWithInformation(registrationObjects.Records)
	_, err = UsersCollection.InsertOne(Ctx, User{
		Name:     registrationObjects.Name,
		Email:    registrationObjects.Email,
		Password: registrationObjects.Password,
		Vector:   vector,
	})
	if err != nil {
		return
	}
	response := InformationVectorResponse{
		Vector:      vector,
		Information: information,
	}
	c.JSON(200, response)
}
