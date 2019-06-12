package models

import (
	"team-members/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Member struct for members
type Member struct {
	OID  *primitive.ObjectID `json:"oid" bson:"_id,omitempty"`
	Name string              `json:"name" bson:"name"`
	Type Type                `json:"type" bson:"type"`
	Tags []string            `json:"tags" bson:"tags"`
}

// Type struct for member type
type Type struct {
	Key        int               `json:"key" bson:"key"`
	Value      string            `json:"value" bson:"value"`
	Properties map[string]string `json:"properties" bson:"properties"`
}

// CleanTypeValue overwrites member type value with appropriate
// type value by type key (in case payload value is incorrect)
func (m Member) CleanTypeValue() Member {

	value := enums.Type.Value(enums.Type(m.Type.Key))

	m.Type.Value = value
	return m
}

// CleanPropertyName overwrites member type property name with appropriate
// type property name by type key (in case payload name is incorrect)
func (m Member) CleanPropertyName() Member {

	name := enums.Property.Name(enums.Property(m.Type.Key))

	var value string
	for _, v := range m.Type.Properties {
		value = v
	}

	property := map[string]string{name: value}
	m.Type.Properties = property
	return m
}
