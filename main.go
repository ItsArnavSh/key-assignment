package main

import (
	"stack/src/server"
)

func main() {
	//codegen.GenerateLargeCodebase()
	server := server.NewServer()
	server.CurrentURL = "https://github.com/go-gorm/gorm/blob/master/clause/clause.go"
	server.Scan("gen")

}
