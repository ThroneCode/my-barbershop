package auth

import (
	"github.com/thronecode/my-barbershop/backend/internal/config/database"
	"github.com/thronecode/my-barbershop/backend/internal/domain/admin"
	"github.com/thronecode/my-barbershop/backend/internal/sorry"
	"github.com/thronecode/my-barbershop/backend/internal/utils"

	"net/http"
)

// Login is the function that logs the user in
func Login(credentials *LoginInput) (*LoginOutput, error) {
	tx, err := database.NewTransaction(true)
	if err != nil {
		return nil, sorry.Err(err)
	}
	defer tx.Rollback()

	adminRepo := admin.New(tx)

	adm, err := adminRepo.GetByUsername(credentials.Username)
	if err != nil {
		return nil, sorry.Err(err)
	}

	if adm == nil || !utils.CheckPassword(*adm.Password, *credentials.Password) {
		return nil, sorry.NewErr("invalid credentials", http.StatusUnauthorized)
	}

	token, err := utils.GenerateToken(*adm.Username, *adm.ID)
	if err != nil {
		return nil, sorry.Err(err)
	}

	return &LoginOutput{Token: &token}, nil
}