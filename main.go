package main

import (
	"femas66/simple-grpc/model"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
)

// main untuk proto saja ;v

var user1 = &model.User{
	Id:       "1",
	Name:     "Femas",
	Password: "1234",
	Gender:   model.UserGender_MALE,
}

var userList = &model.UserList{
	List: []*model.User{user1},
}

func main() {
	fmt.Println(user1)
	fmt.Println(user1.String())

	jsonDataUser, err := protojson.Marshal(user1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonDataUser))
	jsonDataList, err := protojson.Marshal(userList)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonDataList))
	fmt.Println(jsonDataList)

	protoObj := new(model.User)
	err = protojson.Unmarshal(jsonDataUser, protoObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(protoObj)
}
