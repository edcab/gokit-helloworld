package helloworld

import "context"

type HelloInterfaceService interface {
	DoHello(ctx context.Context) (string, error)
}

type HelloService struct {
}

func NewService() HelloService {
	return HelloService{}
}

func (service *HelloService) DoHello(ctx context.Context) (string, error) {
	return "Hola, es tu nuevo amigo Gopher !", nil
}
