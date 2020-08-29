package v1

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1 "github.com/103cuong/go_grpc/pkg/api/v1"
)

const apiVersion = "v1"

type ToDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) *ToDoServiceServer {
	return &ToDoServiceServer{db: db}
}

func (s *ToDoServiceServer) checkAPIVersion(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(
				codes.Unimplemented,
				"unsupported API version: server implements API version '%s', but asked for '%s'", apiVersion, api,
			)
		}
	}
	return nil
}

func (s *ToDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	db, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "database connection failed")
	}
	return db, nil
}

func (s *ToDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	// checking API version
	if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format")
	}

	// inserting into ToDo
	query := "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?,?,?)"
	result, err := db.ExecContext(ctx, query,
		req.ToDo.Title, req.ToDo.Description, reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "inserting into Todo failed")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, status.Error(codes.Unknown, "getting last insert id failed")
	}

	return &v1.CreateResponse{
		Api: apiVersion,
		Id:  id,
	}, nil
}

func (s *ToDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	// checking API version
	if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	// query ToDo by ID
	query := "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?"
	rows, err := db.QueryContext(ctx, query, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "getting element from ToDo failed")
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "retrieving data from Todo failed")
		}
		return nil, status.Error(codes.NotFound, "not found ToDo with ID")
	}

	var todo v1.ToDo
	var reminder time.Time
	if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
		return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
	}

	todo.Reminder, err = ptypes.TimestampProto(reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "reminder field has invalid format")
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, "founding multiple ToDo with ID")
	}

	return &v1.ReadResponse{
		Api:  apiVersion,
		ToDo: &todo,
	}, nil
}

func (s *ToDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	// checking API version
	if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format")
	}

	query := "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?"
	result, err := db.ExecContext(ctx, query,
		req.ToDo.Title, req.ToDo.Description, reminder, req.ToDo.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updating ToDo failed")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "retrieving data from rows affected failed")
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, "ToDo with ID not found")
	}

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: rows,
	}, nil
}

func (s *ToDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	// checking API version
	if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	query := "DELETE FROM ToDo WHERE `ID`=?"
	result, err := db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "deleting ToDo with ID failed")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, status.Error(codes.Unknown, "retrieving data from rows affected failed")
	}

	if rows == 0 {
		return nil, status.Error(codes.NotFound, "ToDo with ID not found")
	}

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: rows,
	}, nil
}

func (s *ToDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	// checking API version
	if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}

	// database connection
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	// closing connection
	defer db.Close()

	query := "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, status.Error(codes.Unknown, "selecting ToDos failed")
	}
	defer rows.Close()

	var reminder time.Time
	var list []*v1.ToDo
	for rows.Next() {
		todo := new(v1.ToDo)
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &reminder); err != nil {
			return nil, status.Errorf(codes.Unknown, "scanning data from Todo failed: %v", err)
		}
		todo.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.Unknown, "reminder field has invalid format")
		}
		list = append(list, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "selecting data failed")
	}

	return &v1.ReadAllResponse{
		Api:   apiVersion,
		ToDos: list,
	}, nil
}
