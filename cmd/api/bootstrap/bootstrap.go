package bootstrap

import (
	"context"
	"fmt"
	"strings"

	"Github.com/NaujOyamat/microservice-template/internal/crosscutting"
	"Github.com/NaujOyamat/microservice-template/internal/infrastructure/httpserver"
	mongoDb "Github.com/NaujOyamat/microservice-template/internal/infrastructure/storage/mongo"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := httpserver.New(host, port, createDependencies())
	return srv.Run()
}

func createDependencies() httpserver.Dependencies {
	db, err := initializeMongoDb()
	if err != nil {
		panic(err)
	}

	return httpserver.Dependencies{
		CourseRepo: mongoDb.NewCourseRepository(db),
	}
}

// //TODO: Update to Go 1.8
// func initializeDb[Db *mongo.Database | *sql.DB]() (Db, error) {
// 	if crosscutting.AppSettings.DbSetting.Provider == crosscutting.MongoProvider {
// 		return initializeMongoDb()
// 	} else {
// 		return initializePostgresDb()
// 	}
// }

func initializeMongoDb() (*mongo.Database, error) {
	userPass := ""
	if strings.Trim(crosscutting.AppSettings.DbSetting.User, " ") != "" {
		userPass = fmt.Sprintf("%s:%s@", crosscutting.AppSettings.DbSetting.User, crosscutting.AppSettings.DbSetting.Password)
	}
	uri := fmt.Sprintf("mongodb://%s%s:%d/?maxPoolSize=20&w=majority",
		userPass, crosscutting.AppSettings.DbSetting.Host, crosscutting.AppSettings.DbSetting.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client.Database("Codely"), nil
}

// func initializePostgresDb() (*sql.DB, error) {
// 	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "Rider170318", "postgres")
// 	return sql.Open("postgres", connStr)
// }
