package pack

import (
	"github.com/lukanzx/DouServer/cmd/interaction/dal/db"
	"github.com/lukanzx/DouServer/kitex_gen/interaction"
	"github.com/lukanzx/DouServer/kitex_gen/user"
)

func Comment(data *db.Comment, user *user.User) *interaction.Comment {
	if data == nil {
		return nil
	}

	return &interaction.Comment{
		Id:         data.Id,
		User:       user,
		Content:    data.Content,
		CreateDate: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
