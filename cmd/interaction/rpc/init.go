package rpc

import (
	"github.com/lukanzx/DouServer/kitex_gen/user/userservice"
	"github.com/lukanzx/DouServer/kitex_gen/video/videoservice"
)

var (
	userClient  userservice.Client
	videoClient videoservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
}
