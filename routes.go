package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

const (
	UrlPong = "/pong"
	UrlFib  = "/fib"
	UrlDb   = "/db"
)

const (
	NonPositiveErr = "number should be a positive number, but it is negative"
	MaxValueErr    = "current value exceeds the limit"
)

const (
	HTTPGet HTTPType = iota
	HTTPPost
	HTTPPut
	HTTPDelete
)

type HTTPType int

type RouteDescriptor struct {
	Path       string
	Method     func(C *gin.Context)
	MethodType HTTPType
}

func Pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Fib(c *gin.Context) {

	number := c.DefaultQuery("number", "0")
	n, err := strconv.Atoi(number)
	if err != nil {
		c.String(
			http.StatusBadRequest, err.Error())
		return
	}
	if n < 0 {
		c.String(http.StatusBadRequest, NonPositiveErr)
		return
	}
	if result := Fibonacci(n); result == -1 {
		c.String(http.StatusInternalServerError, MaxValueErr)
	} else {
		c.String(http.StatusOK, fmt.Sprintf("Result is %d", result))
	}

}

func Db(c *gin.Context) {
	key := c.DefaultQuery("key", "record")
	value := c.DefaultQuery("value", "record_value")

	fmt.Println(key, value)

	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//init
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todo"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return nil
	})
	//write new entry
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		return b.Put([]byte(key), []byte(value))
	})

	//read from DB
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todo"))
		v := b.Get([]byte(key))
		c.String(http.StatusOK, fmt.Sprintf("The value field for '%s' is: %s\n", key, v))
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
