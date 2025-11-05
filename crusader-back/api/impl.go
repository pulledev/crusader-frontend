package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pulledev/crusader-frontend/crusader-back/initializers"
	"github.com/pulledev/crusader-frontend/crusader-back/model"
	"gorm.io/gorm"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

var db *gorm.DB = initializers.GetDB()

type Server struct{}

func jsonToStruct[T any](r *http.Request) *T {
	var obj T
	decoder := json.NewDecoder(r.Body)
	//Good to know: Es wird sehr wahrscheinlich nicht zu err kommen, da alles fehlerhafte bereits in der middleware abgefangen wird. Trotzdem um sicher zu gehen :)
	err := decoder.Decode(&obj)
	if err != nil {
		log.Println("failed to decode Body: ", err)
	}

	return &obj
}

type ErrorResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	GormMessage string `json:"database_message"`
}

func writeErrorInJson(w http.ResponseWriter, status int, msg string, gormErr error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var g string = fmt.Sprint(gormErr)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:        status,
		Message:     msg,
		GormMessage: g,
	})
}

func ptrToUint[T ~int | ~uint | ~float64](p *T) uint {
	if p == nil {
		return 0
	}
	return uint(*p)
}

func (s Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	member := jsonToStruct[MemberCreate](r)

	//DB Shit:

	m := model.Member{
		DiscordId:        member.DiscordId,
		Name:             member.Name,
		SteamId:          member.SteamId,
		UnitID:           *member.Unit,
		MembershipTypeID: member.MembershipType,
		RankLevel:        *member.Rank,
		Stab:             nil,
		DiscordNick:      member.DiscordNick,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}

	result := db.Create(&m)
	if result.Error != nil {

		var msg string = fmt.Sprint("Could not create Member: ", result.Error)
		log.Println(msg)
		fmt.Println(result.Error)

		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			result.
				writeErrorInJson(w, 409, "Mitglied existiert bereits", gorm.ErrDuplicatedKey)

		} else {
			writeErrorInJson(w, 409, "Mitglied konnte nicht erstellt werden", result.Error)

		}

	}

	log.Println("Queried Member:", m)

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
