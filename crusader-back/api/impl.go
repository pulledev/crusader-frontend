package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
	"github.com/pulledev/crusader-frontend/crusader-back/model"
	"gorm.io/gorm"
	"net/http"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

var db *gorm.DB = initializers.GetDB()

type Server struct{}

func (s Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	var member MemberCreate

	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&member)

}

func (s Server) PartialUpdateMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}

func (s Server) RemoveMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}

func (s Server) GetMember(w http.ResponseWriter, r *http.Request, id Id) {
	ctx := context.Background()
	fmt.Println("test")

	member, err := gorm.G[model.Member](db).Where("discord_id = ?", id).First(ctx)
	errors.Is(err, gorm.ErrRecordNotFound)

	if err == nil {

		bites, _ := json.Marshal(member)
		w.Header().Add("Content-Type", "application/json")
		w.Write(bites)
	} else {
		w.WriteHeader(404)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("no Member found"))
	}

}

func (s Server) UpdateMember(w http.ResponseWriter, r *http.Request, id Id) {
	//TODO implement me
	panic("implement me")
}
