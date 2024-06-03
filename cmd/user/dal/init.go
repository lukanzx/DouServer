package dal

import (
	"github.com/lukanzx/DouServer/cmd/user/dal/cache"
	"github.com/lukanzx/DouServer/cmd/user/dal/db"
	"github.com/lukanzx/DouServer/cmd/user/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
