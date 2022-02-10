package service

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/utils"
)

func (u *UserService) AuthUser(email string, password string) (int, error) {
	userDb, err := u.repo.AppUser.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}
	if utils.CheckPasswordHash(password, userDb.Password) {
		return userDb.ID, nil
	} else {
		u.logger.Warn("AuthUser: wrong email or password entered")
		return 0, fmt.Errorf("wrong email or password entered")
	}
}
