package utils

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Port returns port number if environment variable
// is set, otherwise returns default 8080
func Port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

// LogFatal checks if error != nil and if so log.Fatal(err).
//  Fatal is equivalent to Print() followed by a call to os.Exit(1).
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// IsValidOID checks if string is a valid ObjectID hex string length (24)
// and can be converted to ObjectID. If so, returns a new ObjectID created
// from the oid hex string.
func IsValidOID(s string) (bool, primitive.ObjectID) {

	oid, err := primitive.ObjectIDFromHex(s)
	if len(s) == 24 || err == nil {
		return true, oid
	}

	return false, oid
}
