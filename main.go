package main

import (
	"net/http"
	"team-members/controllers"
	"team-members/driver"
	f "team-members/filters"
	"team-members/models"
	"team-members/utils"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	mongo := driver.ConnectDB()

	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// the hello message endpoint with JSON response from map
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework."})
	})

	// get all members
	engine.GET("/api/members", func(c *gin.Context) {

		c.JSON(http.StatusOK, controllers.GetAllMembers(mongo))
	})

	// get members by tags
	engine.GET("/api/members/", func(c *gin.Context) {

		t := c.QueryArray("tags")
		m := controllers.GetMembersByTags(t, mongo)
		c.JSON(http.StatusOK, m)
	})

	// create new member
	engine.POST("/api/member", func(c *gin.Context) {

		var m models.Member
		if err := c.BindJSON(&m); err != nil {

			c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{"message": "Invalid JSON payload."})
		} else {

			i, created := controllers.CreateMember(m, mongo)
			if created {
				c.JSON(http.StatusCreated, gin.H{"message": "Member " + m.Name + " created!"})
			} else {
				switch i {
				case 0:
					c.JSON(http.StatusConflict, gin.H{"message": "Member already exists."})
				case -1:
					c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid type key."})
				}
			}
		}
	})

	//get member by name
	engine.GET("/api/member/name/:name", func(c *gin.Context) {

		n := c.Params.ByName("name")
		m, found := controllers.GetMember(f.Name{Value: n}.Filter(), mongo)
		if found {
			c.JSON(http.StatusOK, m)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Member not found."})
		}
	})

	// get member by ObjectID (:oid param is 24 character ObjectID string)
	engine.GET("/api/member/oid/:oid", func(c *gin.Context) {

		validOID, oid := utils.IsValidOID(c.Params.ByName("oid"))
		if !validOID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID string."})

		} else {
			m, found := controllers.GetMember(f.OID{Value: &oid}.Filter(), mongo)

			if found {
				c.JSON(http.StatusOK, m)
			} else {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Member not found."})
			}
		}
	})

	// update existing member by ObjectID (:oid param is 24 character ObjectID string)
	engine.PUT("/api/member/:oid", func(c *gin.Context) {

		var m models.Member

		validOID, oid := utils.IsValidOID(c.Params.ByName("oid"))
		if !validOID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID string."})

		} else if err := c.BindJSON(&m); err != nil {
			c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{"message": "Invalid JSON payload."})

		} else if m.OID.String() != oid.String() {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Route ObjectID and payload oid are not equal."})

		} else {

			i, updated := controllers.UpdateMember(m, mongo)
			if updated {
				c.JSON(http.StatusOK, gin.H{"message": "Member updated!"})
			} else {
				switch i {
				case 0:
					c.JSON(http.StatusNotFound, gin.H{"message": "Member doesn't exist."})
				case 1:
					c.JSON(http.StatusTeapot, gin.H{"message": "No changes, nothing to update."})
				case -1:
					c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid type key."})
				}
			}
		}
	})

	// delete member
	engine.DELETE("/api/member/:oid", func(c *gin.Context) {

		validOID, oid := utils.IsValidOID(c.Params.ByName("oid"))
		if !validOID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID string."})

		} else {
			deleted := controllers.DeleteMember(&oid, mongo)

			if deleted {
				c.JSON(http.StatusOK, gin.H{"message": "Member deleted!"})
			} else {
				c.JSON(http.StatusNotFound, gin.H{"message": "Member not found, nothing to delete."})
			}
		}
	})

	// run server on PORT
	engine.Run(utils.Port())

}
