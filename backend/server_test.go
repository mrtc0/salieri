package main

import (
	"log"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

type TestPostCode struct {
	Code        string `json:"Code"`
	Language    string `json:"Language"`
	Stdin       string `json:"Stdin"`
	Stdout      string `json:"Stdout"`
	Stderr      string `json:"Stderr"`
	Status_code string `json:"Status_code"`
}

func TestCompilerRequest(t *testing.T) {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/api/compiler/", func(w rest.ResponseWriter, r *rest.Request) {
			code := TestPostCode{}
			code.Code = "#include <stdio.h>\nint main() {\n    printf(\"HELLO\\n\");\n    return 0;\n}\n"
			code.Language = "clang"
			code.Stdin = ""
			code.Stdout = ""
			code.Status_code = ""
			w.WriteJson(&code)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	recorded := test.RunRequest(t, api.MakeHandler(),
		test.MakeSimpleRequest("POST", "http://127.0.0.1/api/compiler/", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}
