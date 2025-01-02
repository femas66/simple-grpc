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

var course1 = &model.Course{
	Id:   "1",
	Name: "Kelas",
	Classroom: &model.Classroom{
		Id:   "1",
		Name: "Kelas 1",
	},
}

var courseList = &model.CourseList{
	List: []*model.Course{course1},
}

var courseListByUser = &model.CourseListByUser{
	List: map[string]*model.CourseList{
		user1.Id: courseList,
	},
}

func main() {
	fmt.Println(user1)
	fmt.Println(user1.String())
	fmt.Println(courseListByUser)

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
