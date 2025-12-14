package main

import (
	initializers2 "github.com/pulledev/crusader-frontend/crusader-back/internal/initializers"
	"github.com/pulledev/crusader-frontend/crusader-back/internal/model"
)

func init() {
	initializers2.LoadEnvVars()
	initializers2.GetDB()

}

func main() {
	initializers2.GetDB().AutoMigrate(&model.Member{})

}
