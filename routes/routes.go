package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "tsacodingchallenge/models"
    "tsacodingchallenge/controllers"
)

// InitializeRoutes sets up the API routes
func InitializeRoutes(router *gin.Engine, client *mongo.Client) {
    router.POST("/contacts", func(c *gin.Context) {
        var contact models.Contact
        if err := c.ShouldBindJSON(&contact); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Validate phone numbers and format them
        for i, phone := range contact.PhoneNumbers {
            formattedPhone, err := controllers.ValidateAndFormatPhone(phone) // Use the correct function
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number: " + err.Error()})
                return
            }
            contact.PhoneNumbers[i] = formattedPhone
        }

        // Insert contact into MongoDB
        collection := client.Database("TSACodingChallenge").Collection("contacts")
        result, err := collection.InsertOne(c, contact)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save contact"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"id": result.InsertedID})
    })

    router.GET("/contacts/:id", func(c *gin.Context) {
        id := c.Param("id")
        objectId, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        // Retrieve contact from MongoDB
        var contact models.Contact
        err = client.Database("TSACodingChallenge").Collection("contacts").FindOne(c, bson.M{"_id": objectId}).Decode(&contact)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
            return
        }

        c.JSON(http.StatusOK, contact)
    })
}
