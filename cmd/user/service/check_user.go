package service

import (
	"github.com/lukanzx/DouServer/cmd/user/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/user"
	"github.com/lukanzx/DouServer/pkg/errno"
	"golang.org/x/crypto/bcrypt"
)

// CheckUser check user is exist and it's password
func (s *UserService) CheckUser(req *user.LoginRequest) (*db.User, error) {
	userModel, err := db.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)) != nil {
		return nil, errno.AuthorizationFailedError
	}

	return userModel, nil
}
