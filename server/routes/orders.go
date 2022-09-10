package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"server/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var orderCollection *mongo.Collection = OpenCollection(Client, "orders")

// add order to orders
func AddOrder(c *gin.Context) {
	// context.WithTimeout pass a context with a timeout to tell a blocking function that it should abandon its work after time timeout elapses
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	validationErr := validate.Struct(order)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		defer cancel()
		return
	}

	order.ID = primitive.NewObjectID()

	result, insertErr := orderCollection.InsertOne(ctx, order)
	if insertErr != nil {
		msg := "order item was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		defer cancel()
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result)
}

// get all orders
func GetOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var orders []bson.M

	cursor, err := orderCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()
	fmt.Println(orders)
	c.JSON(http.StatusOK, orders)
}

// get all orders by the waiter's name
func GetOrdersByWaiter(c *gin.Context) {
	waiter := c.Params.ByName("waiter")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var orders []bson.M

	cursor, err := orderCollection.Find(ctx, bson.M{"server": waiter})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	if err = cursor.All(ctx, &orders); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()
	fmt.Println(orders)
	c.JSON(http.StatusOK, orders)
}

// get an order by its id
func GetOrderById(c *gin.Context) {
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var order bson.M

	if err := orderCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()
	fmt.Println(order)
	c.JSON(http.StatusOK, order)
}

// update a waiter's name for an order
func UpdateWaiter(c *gin.Context) {
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	type Waiter struct {
		Server *string `json:"server"`
	}

	var waiter Waiter

	if err := c.BindJSON(&waiter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	result, err := orderCollection.UpdateOne(
		ctx,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{{Key: "server", Value: waiter.Server}}},
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.MatchedCount)
}

// update an order
func UpdateOrder(c *gin.Context) {
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	if validationErr := validate.Struct(order); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		defer cancel()
		return
	}

	result, err := orderCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"dish":   order.Dish,
			"price":  order.Price,
			"server": order.Server,
			"table":  order.Table,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.ModifiedCount)
}

// delete an order
func DeleteOrder(c *gin.Context) {
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	result, err := orderCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		defer cancel()
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
