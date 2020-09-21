package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/danryan/env"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/unrolled/render"

	"github.com/jooho/localbuddy-server/model"
)

// Config return configuration for server
type Config struct {
	Environment string `env:"key=ENVIRONMENT default=development"`
	Port        string `env:"key=PORT default=9000"`
	EnableCors  string `env:"key=ENABLE_PORT default=false"`
}

var (
	renderer *render.Render
	config   *Config
)

func init() {
	var option render.Options
	config = &Config{}

	if err := env.Process(config); err != nil {
		fmt.Println(err)
	}

	if config.Environment == "development" {
		option.IndentJSON = true
		option.IsDevelopment = true
	}

	renderer = render.New(option)

}

// func renderErrorJson(w http.ResponseWriter, err *delibird.ApiError, status int) {
// 	if status < 100 {
// 		status = http.StatusBadRequest
// 	}

// 	renderer.JSON(w, status, map[string]string{"code": err.Code, "message": err.Message})
// }

func App() http.Handler {

	// router
	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthzRoute)

	r.HandleFunc("/info/{group}", infoGroupPutRoute).Methods("PUT")

	// middleware
	n := negroni.Classic()

	// enable CORS
	if config.EnableCors == "true" {
		c := cors.New(cors.Options{})
		n.Use(c)
	}

	// add handler
	n.UseHandler(r)

	return n

}

// Router Handlers

func healthzRoute(w http.ResponseWriter, r *http.Request) {
	renderer.Text(w, http.StatusOK, "pong")
}

func infoGroupPutRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := []string{"Tax", "Benefit", "Immigration", "HouseAndMortgage", "News"}
	Find(a, vars["groups"])
	// _, found := Find(model.GetGroupTypes, vars["group"])
	_, found := Find(a, vars["group"])

	if !found {

		// renderErrorJson(w, "Group is not exist", 400)
		return
	}

	data := ""
	renderer.JSON(w, http.StatusOK, data)

}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), App())
}
