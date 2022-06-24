package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var entryCollection *mongo.Collection = openColletion(Client, "calories")

func AddEntry(c *gin.Context) {}
func Getentries(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M

	cursor, err := entryCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	// get all from databse Cursor and add to empty slice entries
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entries)
	//Returning all entries in JSON
	c.JSON(http.StatusOK, entries)

}
func GetentryById(c *gin.Context)           {}
func GetEntriesByIngredient(c *gin.Context) {}
func UpdateEntry(c *gin.Context)            {}
func UpdateIngredient(c *gin.Context)       {}

//Function to Delete entry
func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	//To Get the Collection from mongodb and delete: DeleteOne is Default Function in Golang
	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	//Return all the Deleted Count in form of JSON
	c.JSON(http.StatusOK, result.DeletedCount)
}
