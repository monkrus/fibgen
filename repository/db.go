package main

type Bucket string

const (
	result Bucket = "result"
)

type Repository interface {
	Save(n int, result int) error
	Get(n int) (int, error)
}
