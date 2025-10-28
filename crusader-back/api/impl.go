package api

import (
	"fmt"
	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
	"gorm.io/gorm"
	"net/http"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func (s Server) ListMembers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) PostMemberSignin(w http.ResponseWriter, r *http.Request) {
	var signinRequest SigninRequest

	fmt.Println(signinRequest)

	panic("implement me")
}

var db *gorm.DB = initializers.GetDB()
