package main

import (
	"log"
	"proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatalln(srv.ListenAndServe())
	// curl -X POST localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzEK"}}'
	// curl -X GET localhost:8080 -d '{"offset":0}'

	// curl -X POST localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzIK"}}'
	// curl -X GET localhost:8080 -d '{"offset":1}'

	// curl -X POST localhost:8080 -d '{"record":{"value":"TGV0J3MgR28gIzMK"}}'
	// curl -X GET localhost:8080 -d '{"offset":2}'
}
