package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/Anirudh-rao/CalorieTracker-Go-React/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection(Client, "calories")

//Add Entry Function
func AddEntry(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry
	//We are Converting JSON data into Entry Data
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	//Validation of Data
	ValidationErr := validate.Struct(entry)
	if ValidationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ValidationErr.Error()})
		fmt.Println(ValidationErr)
		return
	}
	//Inserting our Converted Data into Collection
	entry.ID = primitive.NewObjectID()
	result, insertErr := entryCollection.InsertOne(ctx, entry)
	if insertErr != nil {
		msg := fmt.Sprintf("Order Item not Created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(msg)
		return
	}
	//Canecl Connection
	defer cancel()
	c.JSON(http.StatusOK, result)
}

//Get Entries Function
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

//Get Entries By ID function
func GetentryById(c *gin.Context) {
	EntryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(EntryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry bson.M
	if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(entry)
	c.JSON(http.StatusOK, entry)

}

//Get Entries By Ingredient Function
func GetEntriesByIngredient(c *gin.Context) {
	ingredient := c.Params.ByName("id")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, entries)
}

//Update Entry Function
func UpdateEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var entry models.Entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	ValidationErr := validate.Struct(entry)
	if ValidationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ValidationErr.Error()})
		fmt.Println(ValidationErr)
		return
	}
	result, err := entryCollection.ReplaceOne(ctx, bson.M{"_id": docID},
		//Takes in the parameters we need to update
		bson.M{
			"dish":        entry.Dish,
			"fat":         entry.Fat,
			"ingredients": entry.Ingredients,
			"calories":    entry.Calories,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

//UpdateIngredient Function
func UpdateIngredient(c *gin.Context) {
	entryID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(entryID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	type Ingredient struct {
		Ingredients *string `json:"ingredients"`
	}
	var ingredient Ingredient
	if err := c.BindJSON(&ingredient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
		bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)

}

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
