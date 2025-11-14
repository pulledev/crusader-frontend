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

var db = initializers.GetDB()

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
	var g = fmt.Sprint(gormErr)
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
		UnitRoleID:       member.UnitRoleId,
		MembershipTypeID: member.MembershipTypeId,
		RankLevel:        member.RankLevel,
		DiscordNick:      member.DiscordNick,
		Points:           member.Points,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}

	result := db.Create(&m)
	if result.Error != nil {

		var msg = fmt.Sprint("Could not create Member: ", result.Error)
		log.Println(msg)

		writeErrorInJson(w, 409, "Mitglied konnte nicht erstellt werden", result.Error)

	}

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

	//member, err := gorm.G[model.Member](db).Where("discord_id = ?", id).First(ctx)

	member, err := gorm.G[model.Member](db).
		Preload("UnitRole", func(db gorm.PreloadBuilder) error { return nil }).
		Preload("UnitRole.Unit", func(db gorm.PreloadBuilder) error { return nil }).
		Preload("MembershipType", func(db gorm.PreloadBuilder) error { return nil }).
		Preload("Rank", func(db gorm.PreloadBuilder) error { return nil }).
		Preload("Stab", func(db gorm.PreloadBuilder) error { return nil }).
		Where("discord_id = ?", id).
		First(ctx)

	errors.Is(err, gorm.ErrRecordNotFound)

	membership := MembershipType{
		CreatedAt: member.MembershipType.CreatedAt,
		Id:        int(member.MembershipType.ID),
		Name:      member.MembershipType.Name,
		UpdatedAt: member.MembershipType.UpdatedAt,
	}
	rank := Rank{
		CreatedAt: member.Rank.CreatedAt,
		Level:     member.Rank.Level,
		Name:      member.Rank.Name,
		UpdatedAt: member.Rank.UpdatedAT,
	}
	unit := Unit{
		CreatedAt:     member.UnitRole.Unit.CreatedAt,
		Description:   member.UnitRole.Unit.Description,
		DiscordRoleId: member.UnitRole.Unit.DiscordRoleId,
		Id:            int(member.UnitRole.Unit.ID),
		Name:          member.UnitRole.Unit.Name,
		UpdatedAt:     member.UnitRole.Unit.UpdatedAt,
	}
	role := UnitRole{
		CreatedAt:   member.UnitRole.CreatedAt,
		Description: member.UnitRole.Description,
		Id:          int(member.UnitRole.ID),
		Name:        member.UnitRole.Name,
		Unit:        unit,
		UpdatedAt:   member.UnitRole.UpdatedAt,
	}
	// Build the response struct with proper type conversions
	response := Member{
		CreatedAt:      &member.CreatedAt,
		DiscordId:      member.DiscordId,
		DiscordNick:    member.DiscordNick,
		MembershipType: membership,
		Name:           member.Name,
		Points:         member.Points,
		Rank:           &rank,
		Stab:           nil,
		SteamId:        member.SteamId,
		UnitRole:       &role,
		UpdatedAt:      &member.UpdatedAt,
	}

	if err == nil {

		bites, _ := json.Marshal(response)
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
