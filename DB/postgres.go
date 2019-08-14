package svc

import (
	"context"
	"fmt"

	"github.com/Infoblox-CTO/atlas.status.service_bkup/pkg/pb"
	toolkit_gorm "github.com/infobloxopen/atlas-app-toolkit/gorm"
	"github.com/jinzhu/gorm"
)

type rserver struct{}

func (r *rserver) Create(context.Context, *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	panic("implement me")
}

// NewBasicServer returns an instance of the default server interface
func NewRecordServer() pb.RecordServiceServer {
	return &rserver{}
}

func CreateRecord() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=atlas_status_service password=postgres sslmode=disable")

	if err != nil {
		fmt.Println("error0 ", err)
		return
	}

	ctx := context.Background()

	PushDataWithTransaction(db, ctx)
	PushDataWithoutTransaction(db, ctx)
	GetData(db)

	fmt.Println("Its successful")
}

func GetData(db *gorm.DB) {
	records := []pb.Record{}

	db.Find(&records)

	for _, record := range records {
		fmt.Printf("%v\n", record)
	}
}

func PushDataWithoutTransaction(db *gorm.DB, ctx context.Context) {
	fmt.Println("Trying PushDataWithoutTransaction")

	recordObj := pb.Record{}
	recordObj.EventKey = "eventkey"
	recordObj.ObjectID = "objectIdWithoutT1"
	recordObj.ObjectType = "objtype"
	recordObj.SubObjectID = "subobjId"
	recordObj.SubObjectType = "subobjType"

	ormObj, err := recordObj.ToORM(ctx)
	if err != nil {
		fmt.Println("error1 ", err)
		return
	}

	if err := db.Create(&ormObj); err != nil {
		fmt.Printf("error2 ", err)
	}
	fmt.Println("Finished PushDataWithoutTransaction")
}

func PushDataWithTransaction(db *gorm.DB, ctx context.Context) {
	fmt.Println("Trying PushDataWithTransaction")
	txn := toolkit_gorm.NewTransaction(db)
	ctx = toolkit_gorm.NewContext(ctx, &txn)

	db, err := toolkit_gorm.BeginFromContext(ctx)
	if err != nil {
		fmt.Println("error1 ", err)
		return
	}

	recordObj := pb.Record{}
	recordObj.EventKey = "eventkey"
	recordObj.ObjectID = "objectIdT1"
	recordObj.ObjectType = "objtype"
	recordObj.SubObjectID = "subobjId"
	recordObj.SubObjectType = "subobjType"

	ormObj, err := recordObj.ToORM(ctx)
	if err != nil {
		fmt.Println("error2 ", err)
		return
	}

	if err := db.Create(&ormObj); err != nil {
		fmt.Printf("error3 ", err)
	}

	recordObj.EventKey = "et1"
	recordObj.ObjectID = "objectIdT2"
	recordObj.ObjectType = "objtype"
	recordObj.SubObjectID = "subobjId"
	recordObj.SubObjectType = "subobjType"

	ormObj, err = recordObj.ToORM(ctx)
	if err != nil {
		fmt.Println("error4 ", err)
		return
	}

	if err := db.Create(&ormObj); err != nil {
		fmt.Printf("error5 ", err)
	}

	db.Commit()
	fmt.Println("Finished PushDataWithTransaction")
}
