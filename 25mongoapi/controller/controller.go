package controller

// controllers & helpers
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gabey14/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database Connection string
const connectionString = "mongodb+srv://learncodeonline:ABgeorge14@cluster0.nlhsz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT

var collection *mongo.Collection

// connect with mongoDB
func init() {
	// client option

	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance/reference
	fmt.Println("Collection instance is ready")

}

// NOTE MONGODB helpers - file

// insert 1 record can return *mongo.InsertOneResult

func insertOneMovie(movie model.Netflix) {
	inserted, error := collection.InsertOne(context.Background(), movie)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Inserted 1 movie in db with id :", inserted.InsertedID)
}

// update 1 record

func updateOneMovie(movieId string) {

	id, error := primitive.ObjectIDFromHex(movieId)
	if error != nil {
		log.Fatal(error)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count :", result.ModifiedCount)
}

// delete 1 record

func deleteOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	// [x] print the deletecount properly
	// fmt.Println("Movie got deleted with delete count:", deleteCount)
	fmt.Println("Movie got deleted with delete count:", deleteCount.DeletedCount)

}

// Delete all records from MONGODB

func deleteAllMovie() int64 {
	// D is ordered representation and M is unordered representation
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount

}

// get all movies from database

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies

}

// NOTE  Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)

}

func CreatMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
	}
	//[x] return the object id properly
	movie.ID = primitive.NewObjectID()
	// inserted_id := insertOneMovie(movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
	// json.NewEncoder(w).Encode(inserted_id.InsertedID)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	parans := mux.Vars(r)

	updateOneMovie(parans["id"])
	json.NewEncoder(w).Encode(parans["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
