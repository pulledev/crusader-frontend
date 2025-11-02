package api

import (
	"net/http"

	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
	"gorm.io/gorm"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

var db *gorm.DB = initializers.GetDB()

type Server struct{}

func (s Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (s Server) RemoveMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetMember(w http.ResponseWriter, r *http.Request, id Id) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Bla Bla Bla"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s Server) PartialUpdateMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}

func (s Server) UpdateMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}
