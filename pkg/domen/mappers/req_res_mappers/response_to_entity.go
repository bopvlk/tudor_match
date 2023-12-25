package reqresmappers

import (
	entities "study_marketplace/pkg/domen/models/entities_models"

	"github.com/markbates/goth"
)

func GothToUserToUser(gothUser goth.User) *entities.User {
	return &entities.User{
		Name:     gothUser.Name,
		Email:    gothUser.Email,
		Photo:    gothUser.AvatarURL,
		Verified: true,
		Role:     "user",
	}
}
