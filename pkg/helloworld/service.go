package helloworld

import (
	"context"
	"fmt"
)

type HelloInterfaceService interface {
	SayHelloToMe(ctx context.Context, request interface{}) (string, error)
}

type HelloService struct {
}

func NewService() HelloService {
	return HelloService{}
}

func (service *HelloService) SayHelloToMe(ctx context.Context, request interface{}) (string, error) {
	fmt.Println("El request en service es: ", request)
	return "Hola, es tu nuevo amigo Gopher !", nil
}
