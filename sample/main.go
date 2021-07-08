package main

import (
	_ "golang_study/sample/matchers"
	"golang_study/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
