package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	api "github.com/yagacc/go-sea-port/domain/api/v1"
	"github.com/yagacc/go-sea-port/domain/domain"
	pb "github.com/yagacc/go-sea-port/domain/repository/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type MockPortRepositoryClient struct {
	mock.Mock
}

func (s *MockPortRepositoryClient) Get(ctx context.Context, req *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	args := s.Called(req)
	return args.Get(0).(*pb.GetResponse), args.Error(1)
}

func (s *MockPortRepositoryClient) Save(ctx context.Context, req *pb.SaveRequest, opts ...grpc.CallOption) (*pb.SaveResponse, error) {
	args := s.Called(req)
	return args.Get(0).(*pb.SaveResponse), args.Error(1)
}

func TestGetShouldReturn404WhenNotFound(t *testing.T) {
	//setup mock behaviour
	mockRepository := &MockPortRepositoryClient{}
	mockRepository.On("Get",
		mock.MatchedBy(func(req *pb.GetRequest) bool {
			return req.PortId == "exists"
		})).
		Return(&pb.GetResponse{Port: &domain.Port{Id: "exists"}}, nil)
	mockRepository.On("Get",
		mock.MatchedBy(func(req *pb.GetRequest) bool {
			return req.PortId == "notfound"
		})).
		Return(&pb.GetResponse{Port: nil}, nil)
	mockRepository.On("Get",
		mock.MatchedBy(func(req *pb.GetRequest) bool {
			return req.PortId == "error"
		})).
		Return(&pb.GetResponse{Port: nil}, status.Error(codes.Internal, "downstream error"))
	//for test
	repoClient := &RepositoryClient{PortRepositoryClient: mockRepository}
	testApi := &Api{RepositoryClient: repoClient}
	tests := []struct {
		name    string
		id      string
		errCode codes.Code
	}{
		{"get when exists", "exists", codes.OK},
		{"get when notfound", "notfound", codes.NotFound},
		{"get when downstream error", "error", codes.Internal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := testApi.Get(context.Background(), &api.GetRequest{
				PortId: tt.id,
			})
			if tt.errCode != codes.OK {
				st, _ := status.FromError(err)
				assert.Equal(t, tt.errCode, st.Code())
			} else {
				assert.Equal(t, tt.id, p.Port.Id)
			}
		})
	}
}
