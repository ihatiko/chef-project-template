package transport

import (
	"context"
	"example/internal/features/peoples"
	protoPeople "example/pkg/protoc/peoples"
)

type transport struct {
	service peoples.IService
}

func (t transport) Update(ctx context.Context, request *protoPeople.UpdateRequest) (*protoPeople.UpdateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t transport) Create(ctx context.Context, request *protoPeople.CreateRequest) (*protoPeople.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t transport) Delete(ctx context.Context, request *protoPeople.DeleteRequest) (*protoPeople.DeleteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t transport) GetById(ctx context.Context, request *protoPeople.GetByIdRequest) (*protoPeople.GetByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t transport) GetByPagination(ctx context.Context, request *protoPeople.GetByPaginationRequest) (*protoPeople.GetByPaginationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func New(service peoples.IService) peoples.ITransport {
	return &transport{service: service}
}
