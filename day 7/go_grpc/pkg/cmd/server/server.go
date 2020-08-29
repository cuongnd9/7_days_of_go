package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/103cuong/go_grpc/pkg/protocol/grpc"
	"github.com/103cuong/go_grpc/pkg/service/v1"
	_ "github.com/go-sql-driver/mysql"
)

func RunServer() error {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:103cuong@/grpc_go?parseTime=true")
	if err != nil {
		return fmt.Errorf("opening database failed")
	}
	defer db.Close()

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, "50000")
}
