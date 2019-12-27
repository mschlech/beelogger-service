package dao
//https://github.com/brunaobh/go-mongodb-rest-api-crud/blob/master/app/dao/flights_dao.go
import (
	context "context"
	"github.com/mongodb/mongo-go-driver/bson"
	. "github.com/mschlech/beelogger-service/cmd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/mgo.v2"
	"log"
	"time"
)

type HiveDao struct {
	Server   string
	Database string
}

var client *mongo.Client

const COLLECTION = "hive"
const DATABASE = "mongodb://localhost:27017"
var db *mongo.Database
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(DATABASE)
	client, _ = mongo.Connect(ctx, clientOptions)
	log.Print("connected to mongodb" , client)
}

func (h *HiveDao) Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	log.Print("connected to mongodb" , client)

	//	db = session.DB("hive").C("hive")

}

var hive Hive
// all hive metrics from ONE
func (h *HiveDao) getHiveMetrics(id string) ([]Hive, error) {
	var hives []Hive

	return hives, nil
}

//from all hives
func (h *HiveDao) GetAllHivesMetrics() ([]*Hive, error) {
	var hives []*Hive
	c := client.Database("hive").Collection("hive")
	log.Print("Db -> ", c.Name())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	names,err := c.Database().ListCollectionNames(ctx,nil,nil)
	if err != nil {
		log.Fatal("no collections")
	}
	log.Print("Collection Names : ",  names)

	findOptions := options.Find()
	findOptions.SetLimit(4)
	cursor, _ := c.Find(ctx, bson.D{}, findOptions)
	//	log.Print("all cursor",cursor.All(ctx, &hives))
	for cursor.Next(context.TODO()) {
		var hiveHolder Hive
		err := cursor.Decode(&hiveHolder)
		if err != nil {
			log.Fatal("no values", err)
		}
		hives = append(hives, &hiveHolder)

	}

	defer cursor.Close(context.TODO())
	log.Print("hives in dao " , hives)
	return  hives, nil
}

//todo
func (h *HiveDao) getTemperatureByHiveId(id string) (Hive, error) {

	return Hive{
		ID: primitive.ObjectID{},
	}, nil
}