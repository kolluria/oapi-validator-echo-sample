// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a new pet to the store
	// (POST /pet)
	AddPet(ctx echo.Context) error
	// Update an existing pet
	// (PUT /pet)
	UpdatePet(ctx echo.Context) error
	// Finds Pets by status
	// (GET /pet/findByStatus)
	FindPetsByStatus(ctx echo.Context, params FindPetsByStatusParams) error
	// Finds Pets by tags
	// (GET /pet/findByTags)
	FindPetsByTags(ctx echo.Context, params FindPetsByTagsParams) error
	// Deletes a pet
	// (DELETE /pet/{petId})
	DeletePet(ctx echo.Context, petId int64, params DeletePetParams) error
	// Find pet by ID
	// (GET /pet/{petId})
	GetPetById(ctx echo.Context, petId int64) error
	// Updates a pet in the store with form data
	// (POST /pet/{petId})
	UpdatePetWithForm(ctx echo.Context, petId int64, params UpdatePetWithFormParams) error
	// uploads an image
	// (POST /pet/{petId}/uploadImage)
	UploadFile(ctx echo.Context, petId int64, params UploadFileParams) error
	// Returns pet inventories by status
	// (GET /store/inventory)
	GetInventory(ctx echo.Context) error
	// Place an order for a pet
	// (POST /store/order)
	PlaceOrder(ctx echo.Context) error
	// Delete purchase order by ID
	// (DELETE /store/order/{orderId})
	DeleteOrder(ctx echo.Context, orderId int64) error
	// Find purchase order by ID
	// (GET /store/order/{orderId})
	GetOrderById(ctx echo.Context, orderId int64) error
	// Create user
	// (POST /user)
	CreateUser(ctx echo.Context) error
	// Creates list of users with given input array
	// (POST /user/createWithList)
	CreateUsersWithListInput(ctx echo.Context) error
	// Logs user into the system
	// (GET /user/login)
	LoginUser(ctx echo.Context, params LoginUserParams) error
	// Logs out current logged in user session
	// (GET /user/logout)
	LogoutUser(ctx echo.Context) error
	// Delete user
	// (DELETE /user/{username})
	DeleteUser(ctx echo.Context, username string) error
	// Get user by user name
	// (GET /user/{username})
	GetUserByName(ctx echo.Context, username string) error
	// Update user
	// (PUT /user/{username})
	UpdateUser(ctx echo.Context, username string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// UpdatePet converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePet(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePet(ctx)
	return err
}

// FindPetsByStatus converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetsByStatus(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByStatusParams
	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetsByStatus(ctx, params)
	return err
}

// FindPetsByTags converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetsByTags(ctx echo.Context) error {
	var err error

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByTagsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetsByTags(ctx, params)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeletePetParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for api_key, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "api_key", runtime.ParamLocationHeader, valueList[0], &ApiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter api_key: %s", err))
		}

		params.ApiKey = &ApiKey
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, petId, params)
	return err
}

// GetPetById converts echo context to params.
func (w *ServerInterfaceWrapper) GetPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPetById(ctx, petId)
	return err
}

// UpdatePetWithForm converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePetWithForm(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdatePetWithFormParams
	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", ctx.QueryParams(), &params.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePetWithForm(ctx, petId, params)
	return err
}

// UploadFile converts echo context to params.
func (w *ServerInterfaceWrapper) UploadFile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "petId", runtime.ParamLocationPath, ctx.Param("petId"), &petId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter petId: %s", err))
	}

	ctx.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UploadFileParams
	// ------------- Optional query parameter "additionalMetadata" -------------

	err = runtime.BindQueryParameter("form", true, false, "additionalMetadata", ctx.QueryParams(), &params.AdditionalMetadata)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter additionalMetadata: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UploadFile(ctx, petId, params)
	return err
}

// GetInventory converts echo context to params.
func (w *ServerInterfaceWrapper) GetInventory(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInventory(ctx)
	return err
}

// PlaceOrder converts echo context to params.
func (w *ServerInterfaceWrapper) PlaceOrder(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PlaceOrder(ctx)
	return err
}

// DeleteOrder converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteOrder(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "orderId", runtime.ParamLocationPath, ctx.Param("orderId"), &orderId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter orderId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteOrder(ctx, orderId)
	return err
}

// GetOrderById converts echo context to params.
func (w *ServerInterfaceWrapper) GetOrderById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "orderId", runtime.ParamLocationPath, ctx.Param("orderId"), &orderId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter orderId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOrderById(ctx, orderId)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// CreateUsersWithListInput converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUsersWithListInput(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateUsersWithListInput(ctx)
	return err
}

// LoginUser converts echo context to params.
func (w *ServerInterfaceWrapper) LoginUser(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginUserParams
	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// ------------- Optional query parameter "password" -------------

	err = runtime.BindQueryParameter("form", true, false, "password", ctx.QueryParams(), &params.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter password: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LoginUser(ctx, params)
	return err
}

// LogoutUser converts echo context to params.
func (w *ServerInterfaceWrapper) LogoutUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.LogoutUser(ctx)
	return err
}

// DeleteUser converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteUser(ctx, username)
	return err
}

// GetUserByName converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserByName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserByName(ctx, username)
	return err
}

// UpdateUser converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithLocation("simple", false, "username", runtime.ParamLocationPath, ctx.Param("username"), &username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateUser(ctx, username)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/pet", wrapper.AddPet)
	router.PUT(baseURL+"/pet", wrapper.UpdatePet)
	router.GET(baseURL+"/pet/findByStatus", wrapper.FindPetsByStatus)
	router.GET(baseURL+"/pet/findByTags", wrapper.FindPetsByTags)
	router.DELETE(baseURL+"/pet/:petId", wrapper.DeletePet)
	router.GET(baseURL+"/pet/:petId", wrapper.GetPetById)
	router.POST(baseURL+"/pet/:petId", wrapper.UpdatePetWithForm)
	router.POST(baseURL+"/pet/:petId/uploadImage", wrapper.UploadFile)
	router.GET(baseURL+"/store/inventory", wrapper.GetInventory)
	router.POST(baseURL+"/store/order", wrapper.PlaceOrder)
	router.DELETE(baseURL+"/store/order/:orderId", wrapper.DeleteOrder)
	router.GET(baseURL+"/store/order/:orderId", wrapper.GetOrderById)
	router.POST(baseURL+"/user", wrapper.CreateUser)
	router.POST(baseURL+"/user/createWithList", wrapper.CreateUsersWithListInput)
	router.GET(baseURL+"/user/login", wrapper.LoginUser)
	router.GET(baseURL+"/user/logout", wrapper.LogoutUser)
	router.DELETE(baseURL+"/user/:username", wrapper.DeleteUser)
	router.GET(baseURL+"/user/:username", wrapper.GetUserByName)
	router.PUT(baseURL+"/user/:username", wrapper.UpdateUser)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"kzsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
