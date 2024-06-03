package dal

import (
	"github.com/lukanzx/DouServer/cmd/interaction/dal/cache"
	"github.com/lukanzx/DouServer/cmd/interaction/dal/db"
	"github.com/lukanzx/DouServer/cmd/interaction/dal/sensitive_words"
)

func Init(path string) {
	db.Init()
	cache.Init()
	sensitive_words.Init(path)
}
