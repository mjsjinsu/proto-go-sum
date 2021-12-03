package data

import (
	userpb "github.com/mjsjinsu/proto-go-sum/protos/v1/user"
)

var UserData = []*userpb.UserMessage{
	{
		UserId:      "1",
		Name:        "Rick",
		PhoneNumber: "01088112655",
		Age:         100,
	},
	{
		UserId:      "2",
		Name:        "MA",
		PhoneNumber: "01047720110",
		Age:         200,
	},
}
