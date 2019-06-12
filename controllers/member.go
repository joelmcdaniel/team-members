package controllers

import (
	"team-members/enums"
	f "team-members/filters"
	"team-members/models"
	mr "team-members/repository/member"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllMembers returns a slice of all members
func GetAllMembers(mongo *mongo.Client) []*models.Member {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	m := r.GetMembers(c, f.All{}.Filter())
	return m
}

// GetMembersByTags returns a slice of all members matching tags
func GetMembersByTags(tags []string, mongo *mongo.Client) []*models.Member {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	m := r.GetMembers(c, f.Tags{Value: tags}.Filter())
	return m
}

// GetMember returns a member according to applied filter
func GetMember(filter bson.D, mongo *mongo.Client) (models.Member, bool) {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	m, found := r.GetMember(c, filter)
	return m, found
}

// CreateMember creates a new Member if the member does not already exist
// which is determined by a filter on ObjectID and member Name
func CreateMember(m models.Member, mongo *mongo.Client) (int, bool) {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	// Check if member already exist by ObjectID
	exists := r.MemberExists(c, f.OID{Value: m.OID}.Filter())
	if exists {
		return 0, false
	}

	// Check if member already exist by Name
	exists = r.MemberExists(c, f.Name{Value: m.Name}.Filter())
	if exists {
		return 0, false
	}

	if enums.IsValidTypeKeyValue(m.Type.Key) {

		m = m.CleanTypeValue().CleanPropertyName()

		created := r.CreateMember(c, m)
		return 1, created
	}
	return -1, false
}

// UpdateMember updates an existing member
func UpdateMember(m models.Member, mongo *mongo.Client) (int64, bool) {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	if enums.IsValidTypeKeyValue(m.Type.Key) {

		m = m.CleanTypeValue().CleanPropertyName()

		matched, updated := r.UpdateMember(c, m, f.OID{Value: m.OID}.Filter())
		return matched, updated
	}
	return -1, false
}

// DeleteMember removes a member by ObjectID
func DeleteMember(oid *primitive.ObjectID, mongo *mongo.Client) bool {
	r := mr.MemberRepository{}
	c := r.MembersCollection(mongo)

	deleted := r.DeleteMember(c, f.OID{Value: oid}.Filter())
	return deleted
}
