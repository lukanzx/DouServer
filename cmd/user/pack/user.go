package pack

import (
	"github.com/lukanzx/DouServer/cmd/user/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/user"
)

func User(data *db.User) *user.User {
	if data == nil {
		return nil
	}

	return &user.User{
		Id:              data.Id,
		Name:            data.Username,
		Avatar:          data.Avatar,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
	}
}
