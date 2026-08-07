package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type application struct {
	redis *redis.Client
}

func (app *application) SetCache(ctx context.Context, key string, value interface{}, duration time.Duration) {
	status := app.redis.Set(ctx, key, value, duration)
	fmt.Println("set status: ", status)
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	
	w.Header().Set("Content-Type", "application/json")
	type Person struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	p := Person{
		Name: "Karan aujla",
		Age:  49,
	}

	jsn, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	app.SetCache(ctx, "person", jsn, 0)

	res, err := app.redis.Get(ctx, "person").Result()
	if err != nil {
		fmt.Fprint(w, err)
	}

	fmt.Fprint(w, res)

	// time.Sleep(2 * time.Second)
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	app := &application{
		redis: rdb,
	}

	http.HandleFunc("/", app.homeHandler)

	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
