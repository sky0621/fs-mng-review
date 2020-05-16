package review

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/sky0621/fs-mng-review/models"

	"github.com/sky0621/fs-mng-grpc/pb"
	"google.golang.org/grpc"
)

func StartServer(serverPort string, db *sql.DB) {
	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReviewServer(s, &server{db: db})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

type server struct {
	pb.ReviewServer
	db *sql.DB
}

// 審査機関の一覧を取得
func (s *server) ListFacility(ctx context.Context, req *pb.ListFacilityRequest) (*pb.ListFacilityResponse, error) {
	log.Printf("req:%#v\n", req)
	mFacilities, err := models.Facilities().All(ctx, s.db)
	if err != nil {
		return nil, err
	}
	var pbFacilities []*pb.Facility
	for _, mFacility := range mFacilities {
		pbFacilities = append(pbFacilities, &pb.Facility{
			ID:   mFacility.ID,
			Name: mFacility.Name,
		})
	}
	return &pb.ListFacilityResponse{
		Facilities: pbFacilities,
	}, nil
}

// 特定(プレ)オーダーの審査記録一覧を取得
func (s *server) ListRecord(ctx context.Context, req *pb.ListRecordRequest) (*pb.ListRecordResponse, error) {
	// FIXME:
	log.Printf("req:%#v\n", req)
	return &pb.ListRecordResponse{
		Records: []*pb.Record{
			{ID: "0b139be0-ba4a-45d7-bd05-37df81a5f9b3", Note: "審査メモ１", RecordDatetime: 1589415917},
			{ID: "d5ea2a4b-377b-4e04-8002-9deddcd75ac9", Note: "審査メモ２", RecordDatetime: 1589416160},
			{ID: "887b225b-5945-445c-babd-b0c9f12bd623", Note: "審査メモ３", RecordDatetime: 1589416188},
		},
	}, nil
}
