package dal

import (
	"github.com/lukanzx/DouServer/cmd/follow/dal/cache"
	"github.com/lukanzx/DouServer/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
