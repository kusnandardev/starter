package logging

import (
	"context"
	"fmt"
	"log"
	"time"

	"kusnandartoni/starter/pkg/setting"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mdb *mongo.Client

type mongologModel struct {
	level string
	file  string
	fx    string
	line  int
	user  string
	msg   string
}

func mongoSetup() {
	now := time.Now()
	var err error
	conParam := fmt.Sprintf("%s://%s:%s@%s",
		setting.MongoDBSetting.Type,
		setting.MongoDBSetting.User,
		setting.MongoDBSetting.Password,
		setting.MongoDBSetting.Host)
	ctx := context.TODO()
	log.Printf("1 ==> %v", time.Since(now))

	mongoOpt := options.Client().ApplyURI(conParam)
	log.Printf("2 ==> %v", time.Since(now))

	mdb, err = mongo.Connect(ctx, mongoOpt)
	if err != nil {
		fmt.Println(err)
		log.Fatal("cant connect mongo")
		return
	}
	log.Printf("3 ==> %v", time.Since(now))

	if err := mdb.Ping(ctx, nil); err != nil {
		fmt.Println(err)
		log.Fatal("cant connect mongo")
	}
	// fmt.Println("Mongo Setup is Ready...")

	timeSpent := time.Since(now)
	log.Printf("Config mongo is ready in %v", timeSpent)
	return
}

func (m *mongologModel) save() error {
	collection := mdb.Database("starter").Collection("logs")
	param := bson.M{
		"level": m.level,
		"file":  m.file,
		"fx":    m.fx,
		"line":  m.line,
		"user":  m.user,
		"msg":   m.msg,
		"time":  time.Now(),
	}

	_, err := collection.InsertOne(context.Background(), param)
	if err != nil {
		return err
	}
	return nil
}
