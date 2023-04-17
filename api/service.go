package api

import (
	"github.com/labstack/echo/v4"
)

type service struct{}

func (service) AddPet(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) UpdatePet(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) FindPetsByStatus(c echo.Context, params FindPetsByStatusParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) FindPetsByTags(c echo.Context, params FindPetsByTagsParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) DeletePet(c echo.Context, petId int64, params DeletePetParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) GetPetById(c echo.Context, petId int64) error {
	//TODO implement me
	panic("implement me")
}

func (service) UpdatePetWithForm(c echo.Context, petId int64, params UpdatePetWithFormParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) UploadFile(c echo.Context, petId int64, params UploadFileParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) GetInventory(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) PlaceOrder(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) DeleteOrder(c echo.Context, orderId int64) error {
	//TODO implement me
	panic("implement me")
}

func (service) GetOrderById(c echo.Context, orderId int64) error {
	//TODO implement me
	panic("implement me")
}

func (service) CreateUser(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) CreateUsersWithListInput(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) LoginUser(c echo.Context, params LoginUserParams) error {
	//TODO implement me
	panic("implement me")
}

func (service) LogoutUser(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (service) DeleteUser(c echo.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (service) GetUserByName(c echo.Context, username string) error {
	//TODO implement me
	panic("implement me")
}

func (service) UpdateUser(c echo.Context, username string) error {
	//TODO implement me
	panic("implement me")
}
