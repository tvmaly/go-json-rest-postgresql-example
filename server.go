package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/jackc/pgx"
	"log"
	"net/http"
)

type Api struct {
	DB *pgx.Conn
}

type Location struct {
	Country string
	City    string
	Zipcode string
	Region  string
}

func extractConfig() pgx.ConnConfig {
	var config pgx.ConnConfig

	config.Host = "localhost"
	config.User = "youruser"
	config.Password = "yourpassword"
	config.Database = "yourdatabase"

	return config
}

func (api *Api) InitDB() {
	var err error
	api.DB, err = pgx.Connect(extractConfig())
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
}

func main() {

	api := Api{}
	api.InitDB()

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/api/v1/locations/:city", api.GetLocation},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))

}

func (api *Api) GetLocation(w rest.ResponseWriter, r *rest.Request) {

	city := r.PathParam("city")
	locations := make([]Location, 0)

	rows, err := api.DB.Query("SELECT country,city,zipcode,coalesce(state,county,' ') FROM version1.locations where city = $1", city)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		var country string
		var city string
		var zipcode string
		var region string

		if err := rows.Scan(&country, &city, &zipcode, &region); err != nil {
			log.Fatal(err)
		}
		locations = append(locations, Location{Country: country, City: city, Zipcode: zipcode, Region: region})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.WriteJson(&locations)

}
