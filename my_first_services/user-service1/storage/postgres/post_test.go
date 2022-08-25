package postgres

import (
//	"fmt"
	"reflect"
	"testing"

	pb "github.com/Gorilla-services/user-service1/genproto"
)
func TestUserRepo_Create(t *testing.T){
	tests :=[]struct{
		name string
		input *pb.User
		want *pb.User
		wantErr bool
	}{
	{
		name : "succescase",
		input : &pb.User{
			FirstName: "hello",
			LastName: "Goodbye",
			Posts:  nil,
			//Id : "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
			//UserId: "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
			//Description: "HI",
			//Medias: nil,
		},
		want : &pb.User{
			FirstName: "hello",
			Id: "",
			LastName: "Goodbye",
			Posts: nil,
			//Description: "HI",
			//Medias: nil,
		},
		wantErr : false,

	}}
	for _,tc :=range tests{
   t.Run(tc.name,func(*testing.T){{
	   got ,err :=repo.CreateUser(tc.input)
	   if err!= nil{
		   t.Fatalf("%s:expected : %v,got :%v",tc.name,tc.wantErr,err)
	   }
	  // fmt.Println(got)
	   got.Id = ""
	   //got.UserId = ""
	  // got.Medias = nil
	   if !reflect.DeepEqual(tc.want,got){
		   t.Fatalf("%s:expected : %v,got :%v",tc.name,tc.want,got)
	   }
   }})

	}
}