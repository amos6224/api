package handlers

import (
	"context"
	"encoding/json"
	"enconding/json"
	"fmt"
	"net/http"

	"github.com/amos6224/api/config"
	"github.com/amos6224/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	collection := client.Database("geekmimi").Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.user 
		if err := cursor.Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} 
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var UpdateUser models.User
	if err := json.NewDecoder((r.Body).Decode(&UpdateUser); err != nil {
		http.Error(w, err.Error), http.StatusBadRequest)
		return
	}

	collection := client.Database("geekmimi").Collection("users")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	result, err := collection.UpdateONe(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	collection := client.Database("geekmimi").Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	json.NewEncoder(w).Encode(result)
}