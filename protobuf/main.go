package main

import (
	"fmt"
	"io"
	"os"

	pb "github.com/vrv501/go-repo/example-grpc/generated"
	"google.golang.org/protobuf/proto"
)

func main() {

	// fmt.Println(&pb.Hello{
	// 	Name: "asdasd",
	// })

	// var sa anypb.Any
	// err := anypb.MarshalFrom(&sa, &pb.SociaMedia{}, proto.MarshalOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// user := &pb.User{
	// 	Id:          10,
	// 	Username:    "asdasdas",
	// 	IsActive:    true,
	// 	Password:    []byte{byte('c'), byte('a')},
	// 	Gender:      pb.Gender_GENDER_FEMALE,
	// 	Emails:      []string{"adasdas", "ASdasd"},
	// 	Address:     &pb.Address{Street: "asdasd", City: "asdasdasd", Country: "asdasdasd"},
	// 	Pos:         &pb.User_Coordinate{Latitude: -0.667, Longitude: 13.45},
	// 	Cchannel:    &sa,
	// 	Chhanellev2: &pb.User_Ss{Ss: &pb.SociaMedia{}},
	// }
	// fmt.Println(user)

	// var sm pb.SociaMedia
	// err = anypb.UnmarshalTo(user.Cchannel, &sm, proto.UnmarshalOptions{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("sm", &sm)

	// if user.Cchannel.MessageIs(&sm) {
	// 	err = user.Cchannel.UnmarshalTo(&sm)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println("sm", &sm)
	// }

	// unknown, err := user.Cchannel.UnmarshalNew()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(unknown.ProtoReflect().Descriptor().Name() == sm.ProtoReflect().Descriptor().Name())

	// jres, err := protojson.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(jres))

	// var res pb.User
	// err = protojson.Unmarshal(jres, &res)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(&res)

	// userGrp := &pb.UserGroup{
	// 	GroupId:   10,
	// 	GroupName: "asdasdasd",
	// 	Roles:     []string{"ADasdasdas", "asdasda"},
	// 	Users:     []*pb.User{user},
	// }
	// fmt.Println(userGrp)

	// jres, err = protojson.Marshal(userGrp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(jres))

	// var resGrp pb.UserGroup
	// err = protojson.Unmarshal(jres, &resGrp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(&resGrp)
	// RWV1()
	RWV2()
}

func RWV1() {
	Write(&pb.UserContent{
		UserContentId: 30,
		Slug:          "asdad",
		HtmlContent:   "asdasdas",
		AuthorId:      30,
	})

	Read("any.bin")
}

func RWV2() {

	Read("anyv1.bin")
	Write(&pb.UserContent{
		UserContentId: 30,
		Slug:          "asdasd",
		HtmlContent:   "asdasd",
		AuthorId:      0222,
		Cateogery:     "sdasdasdasd",
	})
	Read("any.bin")
}

func Write(m *pb.UserContent) {
	byte, _ := proto.Marshal(m)
	_ = os.WriteFile("any.bin", byte, 0o644)

}

func Read(fname string) {
	f, _ := os.Open(fname)
	defer f.Close()
	byte, _ := io.ReadAll(f)

	var m pb.UserContent
	_ = proto.Unmarshal(byte, &m)
	fmt.Println(&m)
}
