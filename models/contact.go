package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Contact represents a contact in the application
type Contact struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    FullName    string             `bson:"full_name" json:"full_name"`
    Email       string             `bson:"email" json:"email"`
    PhoneNumbers []string          `bson:"phone_numbers" json:"phone_numbers"`
}
