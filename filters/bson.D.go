package filters

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Filterer interface
type Filterer interface {
	Filter() bson.D
}

// All ...
type All struct{}

// OID ...
type OID struct{ Value *primitive.ObjectID }

// Name ...
type Name struct{ Value string }

// Tags ...
type Tags struct{ Value []string }

// Filter All allows all members returned
func (a All) Filter() bson.D {
	return bson.D{{}}
}

// Filter OID filters on ObjectID
func (oid OID) Filter() bson.D {
	return bson.D{{Key: "_id", Value: oid.Value}}
}

// Filter Name filters on name
func (n Name) Filter() bson.D {
	return bson.D{{Key: "name", Value: n.Value}}
}

// Filter Tags filters on tags
func (t Tags) Filter() bson.D {

	d := bson.D{{
		Key: "tags",
		Value: bson.D{{
			Key:   "$in",
			Value: t.Value,
		}},
	}}

	return d
}
