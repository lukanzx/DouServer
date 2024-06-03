package rpc

import (
	"github.com/lukanzx/DouServer/kitex_gen/chat/messageservice"
	"github.com/lukanzx/DouServer/kitex_gen/follow/followservice"
	"github.com/lukanzx/DouServer/kitex_gen/interaction/interactionservice"
	"github.com/lukanzx/DouServer/kitex_gen/user/userservice"
	"github.com/lukanzx/DouServer/kitex_gen/video/videoservice"
)

var (
	userClient        userservice.Client
	followClient      followservice.Client
	interactionClient interactionservice.Client
	chatClient        messageservice.Client
	videoClient       videoservice.Client
)

func Init() {
	InitUserRPC()
	InitFollowRPC()
	InitInteractionRPC()
	InitChatRPC()
	InitVideoRPC()
}
