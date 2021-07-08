package main

import (
	_ "github.com/Seventeen1Gx/golang_study/sample/matchers"
	"github.com/Seventeen1Gx/golang_study/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
