package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func GetAuth(c *gin.Context) {
	var authorizationObject AuthorizationRequestObject
	var result AuthorizationResponseObject
	err := c.BindJSON(&authorizationObject)
	if err != nil {
		c.Error(err)
	}
	_, vector := CalculateVectorWithInformation(authorizationObject.Records)

	findOptions := options.Find()
	cur, err := UsersCollection.Find(Ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(Ctx) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		isNeededUser := isVectorsEquals(vector, elem.Vector) && elem.Email == authorizationObject.Email

		if isNeededUser {
			result = AuthorizationResponseObject{
				Message: "User founded! Hello " + elem.Name + "!",
				Name:    elem.Name,
			}
			c.JSON(200, result)
			return
		}

	}
	result = AuthorizationResponseObject{
		Message: "User not found",
		Name:    "Name not found",
	}
	c.JSON(200, result)
	return
}
