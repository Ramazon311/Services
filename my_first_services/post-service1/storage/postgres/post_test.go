package postgres

import (
	"testing"
	"reflect"

	pb "github.com/Gorilla-services/post-service1/genproto"
)
func TestPostRepo_Create(t *testing.T){
	tests :=[]struct{
		name string
		input *pb.Post
		want *pb.Post
		wantErr bool
	}{
	{
		name : "succescase",
		input : &pb.Post{
			Name: "hello",
			//Id : "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
			UserId: "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
			Description: "HI",
			Medias: nil,
		},
		want : &pb.Post{
			Name: "hello",
			Id: "",
			UserId: "",
			Description: "HI",
			Medias: nil,
		},
		wantErr : false,

	}}
	for _,tc :=range tests{
   t.Run(tc.name,func(t *testing.T){{
	   got ,err :=repo.Create(tc.input)
	   if err!= nil{
		   t.Fatalf("%s:expected : %v,got :%v",tc.name,tc.wantErr,err)
	   }
	   got.Id = ""
	   got.UserId = ""
	   got.Medias = nil
	   if !reflect.DeepEqual(tc.want,got){
		   t.Fatalf("%s:expected : %v,got :%v",tc.name,tc.want,got)
	   }
   }})

	}
}