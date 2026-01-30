package main

import (
	"encoding/json"
	"fmt"
	"go-kasir-api/database"
	"go-kasir-api/handlers"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"hello": "world",
		})
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllCategories(w, r)
		case "POST":
			handlers.StoreCategory(w, r)
		}
	})

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetCategoryById(w, r)
		case "PUT":
			handlers.UpdateCategory(w, r)
		case "DELETE":
			handlers.DeleteCategory(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API is Running",
		})
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Failed to running server:", err)
	}
}
