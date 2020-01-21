package dao

//https://github.com/brunaobh/go-mongodb-rest-api-crud/blob/master/app/dao/flights_dao.go
import (
	context "context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	. "github.com/mschlech/beelogger-service/cmd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "gopkg.in/mgo.v2"
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
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err!=nil {
		log.Fatal("cannot connect to Database")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client.Connect(ctx)
	log.Print("connected to mongodb", client)
}

//func (h *HiveDao) Connect() {
//	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
//	if err!=nil {
//		log.Fatal("cannot connect to Database")
//	}
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client.Connect(ctx)
//	//client, err := mongo.Connect(ctx, clientOptions)
//	log.Print("error : " , err )
//	log.Print("listDatabaseNames : ", client.Database(DATABASE).Name())
//	log.Print("Ping : ", client.Ping(context.TODO(),nil))
//	log.Print("connected to mongodb", client)
//	log.Print("Name ", client.Database("hive").Name())
//	c := client.Database("hive").Collection("hive")
//	log.Print("Db -> ", c.Name())
//
//	var hives []bson.M
//	cursor, err := c.Find(ctx, bson.M{})
//	if err!=nil {
//		log.Fatal(err)
//	}
//	if err = cursor.All(ctx,&hives); err != nil {
//		log.Fatal(err)
//	}
//	log.Print(hives)
//}

var hive Hive

// all hive metrics from ONE
func (h *HiveDao) getHiveMetrics(id string) ([]Hive, error) {
	var hives []Hive

	return hives, nil
}

//func (h *HiveDao) GetAllHiveMetrics()  {
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	c := client.Database("hive").Collection("hive")
//		var hives []bson.M
//		cursor, err := c.Find(ctx, bson.M{})
//		if err!=nil {
//			log.Fatal(err)
//		}
//		if err = cursor.All(ctx,&hives); err != nil {
//			log.Fatal(err)
//		}
//		log.Print(hives)
//}
//from all hives
func (h *HiveDao) GetAllHivesMetrics() ([]*Hive, error) {
	var hives []*Hive

	c := client.Database("hive").Collection("hive")
	log.Print("Db -> ", c.Name())

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)


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
	log.Print("hives in dao ", hives)
	return hives, nil
}

//todo
func (h *HiveDao) getTemperatureByHiveId(id string) (Hive, error) {

	return Hive{
		ID: primitive.ObjectID{},
	}, nil
}
