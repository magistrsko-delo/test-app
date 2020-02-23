package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"test-app/models"
	"test-app/router"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}


func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		log.Println("middleware", r.URL)
		w.Header().Set("Content-Type", "application/json")
		// can be used for user token hadling.. CORS policy
		h.ServeHTTP(w, r)
	})
}

func main()  {

	envKey := os.Getenv("KEY")
	fmt.Println(envKey)

	models.InitEnv(envKey)
	envStruct := models.GetEnvStruct()

	fmt.Println("-----------------------")
	fmt.Println(envStruct.GetEnv())

	fmt.Println()
	fmt.Println(os.Getenv("KEY"))

	/*env, _ := godotenv.Unmarshal("KEY=whatever")
	fmt.Println(env)
	_ = godotenv.Write(env, "./.env")*/

	fmt.Println("entry go")

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	var customRouter *router.CustomRouter = &router.CustomRouter{R: api}
	customRouter.RegisterHandlers()

	api.Use(Middleware)

	log.Fatal(http.ListenAndServe(":8000", r))
}
