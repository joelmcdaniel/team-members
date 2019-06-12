package memberrepository

import (
	"context"
	"os"
	"team-members/models"
	"team-members/utils"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MemberRepository struct used for calling memberrepository methods
type MemberRepository struct{}

// MembersCollection performs operations on a collection of members
func (r MemberRepository) MembersCollection(m *mongo.Client) *mongo.Collection {
	return m.Database(os.Getenv("DB_NAME")).Collection("members")
}

// GetMembers returns all members
func (r MemberRepository) GetMembers(c *mongo.Collection, filter bson.D) []*models.Member {
	ctx := context.TODO()

	cur, err := c.Find(ctx, filter, options.Find())
	utils.LogFatal(err)

	defer cur.Close(ctx)

	var members []*models.Member

	for cur.Next(ctx) {
		var m models.Member
		err := cur.Decode(&m)
		utils.LogFatal(err)
		members = append(members, &m)
	}
	utils.LogFatal(cur.Err())

	return members
}

// GetMember gets a members according to filter
func (r MemberRepository) GetMember(c *mongo.Collection, filter bson.D) (models.Member, bool) {

	var m models.Member
	err := c.FindOne(context.TODO(), filter).Decode(&m)

	return m, err == nil
}

// MemberExists finds if a member exists according to filter
func (r MemberRepository) MemberExists(c *mongo.Collection, filter bson.D) bool {
	ctx := context.TODO()

	cur, err := c.Find(ctx, filter, options.Find().SetLimit(1))
	utils.LogFatal(err)

	defer cur.Close(ctx)

	exists := cur.Next(ctx)
	utils.LogFatal(cur.Err())

	return exists
}

// CreateMember inserts a new member
func (r MemberRepository) CreateMember(c *mongo.Collection, m models.Member) bool {

	ir, err := c.InsertOne(context.TODO(), m)
	utils.LogFatal(err)

	return ir.InsertedID != nil
}

// UpdateMember updates an existing member
func (r MemberRepository) UpdateMember(c *mongo.Collection, m models.Member, filter bson.D) (int64, bool) {
	u := bson.D{{Key: "$set", Value: m}}

	ur, err := c.UpdateOne(context.TODO(), filter, u)
	utils.LogFatal(err)

	return ur.MatchedCount, ur.ModifiedCount == 1
}

// DeleteMember deletes a member according to filter
func (r MemberRepository) DeleteMember(c *mongo.Collection, filter bson.D) bool {

	dr, err := c.DeleteOne(context.TODO(), filter)
	utils.LogFatal(err)

	return dr.DeletedCount == 1
}
