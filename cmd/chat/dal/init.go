package dal

import (
	"github.com/lukanzx/DouServer/cmd/chat/dal/cache"
	"github.com/lukanzx/DouServer/cmd/chat/dal/db"
	"github.com/lukanzx/DouServer/cmd/chat/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.InitRabbitMQ()
	mq.InitChatMQ()
}
