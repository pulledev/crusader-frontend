package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pulledev/crusader-frontend/crusader-back/internal/initializers"
	"github.com/pulledev/crusader-frontend/crusader-back/internal/model"
	"gorm.io/gorm"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

var db = initializers.GetDB()

type Server struct{}

func (s Server) ListMembers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

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

	//Fügt member hinzu
	result := db.Create(&m)

	if result.Error != nil {
		log.Println(fmt.Sprint("Could not create Member: ", result.Error))
		writeErrorInJson(w, 409, "Mitglied konnte nicht erstellt werden", result.Error)

	}

	stabs := make([]model.Stab, 0, len(*member.StabIds))
	for _, id := range *member.StabIds {
		stabs = append(stabs, model.Stab{
			Model: gorm.Model{ID: uint(id)},
		})
	}

	if err := db.Model(&m).Association("Stab").Append(&stabs); err != nil {
		log.Println(fmt.Sprint("Could not create Member Stab Association: ", result.Error))
		writeErrorInJson(w, 409, "Mitglied Stab Assoziation konnte nicht erstellt werden", result.Error)
	}

}

func (s Server) PartialUpdateMember(w http.ResponseWriter, r *http.Request, id Id) {
	p := jsonToStruct[MemberPartialUpdate](r)

	updates := map[string]any{}

	if p.DiscordNick != nil {
		updates["discord_nick"] = *p.DiscordNick
	}
	if p.MembershipTypeId != nil {
		updates["membership_type_id"] = *p.MembershipTypeId
	}
	if p.Name != nil {
		updates["name"] = *p.Name
	}
	if p.Points != nil {
		updates["points"] = *p.Points
	}
	if p.RankLevel != nil {
		updates["rank_level"] = *p.RankLevel
	} // oder rank_id, je nach Schema
	if p.SteamId != nil {
		updates["steam_id"] = *p.SteamId
	}
	if p.UnitRoleId != nil {
		updates["unit_role_id"] = *p.UnitRoleId
	}

	if len(updates) > 0 {
		if err := db.Model(&model.Member{}).
			Where("discord_id = ?", id).
			Updates(updates).Error; err != nil {
		}
	}

	if p.StabIds != nil {

		var m model.Member

		if err := db.Select("discord_id").Where("discord_id = ?", id).First(&m).Error; err != nil {
			writeErrorInJson(w, http.StatusBadRequest, "Fehler bei select von discord_id", err)
		}

		var stabs []model.Stab
		if len(*p.StabIds) > 0 {
			if err := db.Where("id IN ?", *p.StabIds).Find(&stabs).Error; err != nil {
				writeErrorInJson(w, http.StatusBadRequest, "Fehler bei par update von stab ids", err)
			}
		}

		if err := db.Model(&m).Association("Stab").Replace(stabs); err != nil {
			writeErrorInJson(w, http.StatusBadRequest, "Fehler bei par update von stab ids join", err)
		}
	}

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

	if err != nil {
		writeErrorInJson(w, 409, "Mitglied Stab Assoziation konnte nicht erstellt werden", err)
	}

	membership := MembershipType{
		CreatedAt: member.MembershipType.CreatedAt,
		Id:        int(member.MembershipType.ID),
		Name:      member.MembershipType.Name,
		UpdatedAt: member.MembershipType.UpdatedAt,
	}
	var rank *Rank
	if member.Rank != nil {
		rank = &Rank{
			CreatedAt: member.Rank.CreatedAt,
			Level:     member.Rank.Level,
			Name:      member.Rank.Name,
			UpdatedAt: member.Rank.UpdatedAT,
		}
	}
	var unit *Unit
	if member.UnitRole != nil {
		unit = &Unit{
			CreatedAt:     member.UnitRole.Unit.CreatedAt,
			Description:   member.UnitRole.Unit.Description,
			DiscordRoleId: member.UnitRole.Unit.DiscordRoleId,
			Id:            int(member.UnitRole.Unit.ID),
			Name:          member.UnitRole.Unit.Name,
			UpdatedAt:     member.UnitRole.Unit.UpdatedAt,
		}
	}
	var role *UnitRole
	if member.UnitRole != nil {
		role = &UnitRole{
			CreatedAt:   member.UnitRole.CreatedAt,
			Description: member.UnitRole.Description,
			Id:          int(member.UnitRole.ID),
			Name:        member.UnitRole.Name,
			UpdatedAt:   member.UnitRole.UpdatedAt,
		}
		if unit != nil {
			role.Unit = *unit
		}
	}

	var stab []Stab

	for _, st := range member.Stab {
		stab = append(stab,
			Stab{
				CreatedAt:   st.CreatedAt,
				Description: st.Description,
				Id:          int(st.ID),
				Name:        st.Name,
				UpdatedAt:   st.UpdatedAt,
			})
	}

	response := Member{
		CreatedAt:      &member.CreatedAt,
		DiscordId:      member.DiscordId,
		DiscordNick:    member.DiscordNick,
		MembershipType: membership,
		Name:           member.Name,
		Points:         member.Points,
		Rank:           rank,
		Stab:           &stab,
		SteamId:        member.SteamId,
		UnitRole:       role,
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
	/*member := jsonToStruct[MemberUpdate](r)

	if err == nil {
		bites, _ := json.Marshal(response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(bites)
	} else {
		w.WriteHeader(404)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("no Member found"))
	}
	*/

	panic("implement me")
}
