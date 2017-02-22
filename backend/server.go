package main

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"./core"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/spf13/viper"
)

var bind = "127.0.0.1"
var port = 8080
var name = "jessie2"

func Init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("config/")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("[-] Failed Config Load : %s", err)
	}
	bind = viper.GetString("development.server")
	port = viper.GetInt("development.port")
	name = viper.GetString("development.lxcname")
}

func main() {
	Init()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return origin == "http://localhost:3000"
		},
		AllowedMethods:                []string{"GET", "POST", "PUT"},
		AllowedHeaders:                []string{"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	router, err := rest.MakeRouter(
		rest.Get("/api/compiler", GetAllCompilers),
		rest.Get("/api/compiler/:name", GetCompilerDetails),
		rest.Post("/api/compiler/", Compile),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(bind+":"+strconv.Itoa(port), api.MakeHandler()))
}

type Compiler struct {
	Name    string
	Version string
}

type PostCode struct {
	Code     string
	Language string
	Stdin    string
	Stdout   string
	Stderr   string
}

var store = map[string]*Compiler{}
var postcode_store = map[string]*PostCode{}
var lock = sync.RWMutex{}

func GetAllCompilers(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	compilers := make([]Compiler, len(store))
	i := 0
	for _, compiler := range store {
		compilers[i] = *compiler
		i++
	}
	lock.RUnlock()
	w.WriteJson(&compilers)
}

func GetCompilerDetails(w rest.ResponseWriter, r *rest.Request) {
	// code := r.PathParam("name")
}

func Compile(w rest.ResponseWriter, r *rest.Request) {
	code := PostCode{}
	err := r.DecodeJsonPayload(&code)

	// Error Handling
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if code.Code == "" {
		rest.Error(w, "code required", 400)
		return
	}
	if code.Language == "" {
		rest.Error(w, "language required", 400)
		return
	}

	// Code Push to Container
	err = core.CodePush(name, code.Code, code.Language)
	if err != nil {
		rest.Error(w, "Failed push code", 400)
		return
	}
	// Compile
	result := core.Compile(name, code.Language, code.Stdin)
	code.Stdout = result["stdout"]
	code.Stderr = result["stderr"]

	lock.Lock()
	postcode_store[code.Language] = &code
	lock.Unlock()
	w.WriteJson(&code)

}
