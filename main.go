package main

import (
	"fyne.io/fyne/v2/storage/repository"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"

	log "github.com/sirupsen/logrus"
)

// covers big and negative number scenarios
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n >= 25 {
		return -1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func setupRouter() *MyGinServer {
	gin.SetMode(gin.ReleaseMode)
	r := NewMyGinServer("test name")

	r.InitRoutes([]RouteDescriptor{
		{Path: UrlPong, Method: Pong, MethodType: HTTPGet},
		{Path: UrlFib, Method: Fib, MethodType: HTTPGet},
	})

	return r
}

const bucketName = "todo"

func main() {
	err := setupRouter().Run(":8080")
	if err != nil {
		log.Fatalf("Internal error: %v", err)
	

		//install
		db, err := bolt.Open("bolt.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//init
		db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists("todo")
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		})
        //write new entry
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("todo"))
			return b.Put([]byte("key"), []byte("value"))
		})

		//read from DB
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("todo"))
			v := b.Get([]byte(" key"))
			fmt.Printf("The value field for 'key' is: %s\n", v)
			return nil
		})


	//coloration for logging
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())

	//sample logging
	log.WithFields(log.Fields{
		"user": "admin",
	}).Info("Some info")

	log.Warn("This is a warning")
	log.Error("An error occured!")
}
