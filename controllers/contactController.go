package controllers

import (
    "context"
    "errors"
    "net/http"
    "tsacodingchallenge/models"
    "github.com/gin-gonic/gin"
    "github.com/nyaruka/phonenumbers"
    "go.mongodb.org/mongo-driver/mongo"
)

// ValidateAndFormatPhone validates and formats the phone number to E.164 format.
func ValidateAndFormatPhone(phone string) (string, error) {
    parsedNumber, err := phonenumbers.Parse(phone, "AU")
    if err != nil {
        return "", err
    }

    if !phonenumbers.IsValidNumber(parsedNumber) {
        return "", errors.New("invalid phone number")
    }

    return phonenumbers.Format(parsedNumber, phonenumbers.E164), nil
}

// AddContact adds a new contact to the database.
func AddContact(c *gin.Context, client *mongo.Client, ctx context.Context) {
    var contact models.Contact
    if err := c.ShouldBindJSON(&contact); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    for i, phone := range contact.PhoneNumbers {
        formattedPhone, err := ValidateAndFormatPhone(phone)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        contact.PhoneNumbers[i] = formattedPhone
    }

    collection := client.Database("TSACodingChallenge").Collection("contacts")
    _, err := collection.InsertOne(ctx, contact)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert contact"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Contact added successfully"})
}
