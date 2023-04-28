package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	servergen "github.com/oapi-validator-echo-sample/server-gen"
	"github.com/oapi-validator-echo-sample/utils"
)

func NewEchoServer(ctx context.Context, middlewares ...echo.MiddlewareFunc) *echo.Echo {
	e := echo.New()

	// Register middleware
	for _, middleware := range middlewares {
		e.Use(middleware)
	}

	// Use custom validator and error handler
	e.Validator = utils.NewValidator(ctx)
	e.HTTPErrorHandler = ErrorHandler

	// Register handlers
	servergen.RegisterHandlers(e, servergen.NewStrictHandler(
		newServerInstance(),
		[]servergen.StrictMiddlewareFunc{
			validateStructMiddleware,
		},
	))

	return e
}

// strictServer is the implementation of the servergen.Server interface
type strictServer struct{}

func (s strictServer) AddPet(ctx context.Context, request servergen.AddPetRequestObject) (servergen.AddPetResponseObject, error) {
	fmt.Println("endpoint: AddPet")
	return servergen.AddPet200JSONResponse{}, nil
}

func (s strictServer) UpdatePet(ctx context.Context, request servergen.UpdatePetRequestObject) (servergen.UpdatePetResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) FindPetsByStatus(ctx context.Context, request servergen.FindPetsByStatusRequestObject) (servergen.FindPetsByStatusResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) FindPetsByTags(ctx context.Context, request servergen.FindPetsByTagsRequestObject) (servergen.FindPetsByTagsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) DeletePet(ctx context.Context, request servergen.DeletePetRequestObject) (servergen.DeletePetResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) GetPetById(ctx context.Context, request servergen.GetPetByIdRequestObject) (servergen.GetPetByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) UpdatePetWithForm(ctx context.Context, request servergen.UpdatePetWithFormRequestObject) (servergen.UpdatePetWithFormResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) UploadFile(ctx context.Context, request servergen.UploadFileRequestObject) (servergen.UploadFileResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) GetInventory(ctx context.Context, request servergen.GetInventoryRequestObject) (servergen.GetInventoryResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) PlaceOrder(ctx context.Context, request servergen.PlaceOrderRequestObject) (servergen.PlaceOrderResponseObject, error) {
	return servergen.PlaceOrder200JSONResponse(*request.JSONBody), nil
}

func (s strictServer) DeleteOrder(ctx context.Context, request servergen.DeleteOrderRequestObject) (servergen.DeleteOrderResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) GetOrderById(ctx context.Context, request servergen.GetOrderByIdRequestObject) (servergen.GetOrderByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) CreateUser(ctx context.Context, request servergen.CreateUserRequestObject) (servergen.CreateUserResponseObject, error) {
	fmt.Println("endpoint: CreateUser")
	return servergen.CreateUserdefaultJSONResponse{
		Body:       *request.Body,
		StatusCode: http.StatusOK,
	}, nil
}

func (s strictServer) CreateUsersWithListInput(ctx context.Context, request servergen.CreateUsersWithListInputRequestObject) (servergen.CreateUsersWithListInputResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) LoginUser(ctx context.Context, request servergen.LoginUserRequestObject) (servergen.LoginUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) LogoutUser(ctx context.Context, request servergen.LogoutUserRequestObject) (servergen.LogoutUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) DeleteUser(ctx context.Context, request servergen.DeleteUserRequestObject) (servergen.DeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) GetUserByName(ctx context.Context, request servergen.GetUserByNameRequestObject) (servergen.GetUserByNameResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s strictServer) UpdateUser(ctx context.Context, request servergen.UpdateUserRequestObject) (servergen.UpdateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func newServerInstance() *strictServer {
	return &strictServer{}
}

func validateStructMiddleware(s servergen.StrictHandlerFunc, operationID string) servergen.StrictHandlerFunc {
	return func(ctx echo.Context, i interface{}) (interface{}, error) {
		if err := ctx.Validate(i); err != nil {
			return nil, fmt.Errorf("%s failed to validate request body: %w", operationID, err)
		}

		return s(ctx, i)
	}
}
