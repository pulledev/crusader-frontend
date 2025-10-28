package main

import (
	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
	"github.com/pulledev/crusader-frontend/crusader-back/model"
)

func init() {
	initializers.LoadEnvVars()
	initializers.GetDB()

}

func main() {
	initializers.GetDB().AutoMigrate(&model.Member{})

}
