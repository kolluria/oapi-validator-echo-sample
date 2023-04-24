// Package servergen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package servergen

const (
	Api_keyScopes       = "api_key.Scopes"
	Petstore_authScopes = "petstore_auth.Scopes"
)

// Defines values for ApiResponseType.
const (
	Error   ApiResponseType = "error"
	Success ApiResponseType = "success"
)

// Defines values for OrderStatus.
const (
	Approved  OrderStatus = "approved"
	Delivered OrderStatus = "delivered"
	Placed    OrderStatus = "placed"
)

// Defines values for PetStatus.
const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

// Defines values for FindPetsByStatusParamsStatus.
const (
	FindPetsByStatusParamsStatusAvailable FindPetsByStatusParamsStatus = "available"
	FindPetsByStatusParamsStatusPending   FindPetsByStatusParamsStatus = "pending"
	FindPetsByStatusParamsStatusSold      FindPetsByStatusParamsStatus = "sold"
)

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Code    int32            `json:"code"`
	Message string           `json:"message"`
	Type    *ApiResponseType `json:"type,omitempty" validate:"oneof=success error"`
}

// ApiResponseType defines model for ApiResponse.Type.
type ApiResponseType string

// Category defines model for Category.
type Category struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Order defines model for Order.
type Order struct {
	Complete *bool   `json:"complete,omitempty"`
	Id       *int64  `json:"id,omitempty"`
	PetId    *int64  `json:"petId,omitempty"`
	Quantity *int32  `json:"quantity,omitempty" validate:"gte=1,lte=100"`
	ShipDate *string `json:"shipDate,omitempty" validate:"datetime"`

	// Status Order Status
	Status *OrderStatus `json:"status,omitempty" validate:"oneof=placed approved delivered"`
}

// OrderStatus Order Status
type OrderStatus string

// Pet defines model for Pet.
type Pet struct {
	Category  *Category `json:"category,omitempty"`
	Id        *int64    `json:"id,omitempty"`
	Name      string    `json:"name" validate:"required"`
	PhotoUrls []string  `json:"photoUrls" validate:"required"`

	// Status pet status in the store
	Status *PetStatus `json:"status,omitempty" validate:"oneof=available pending sold"`
	Tags   *[]Tag     `json:"tags,omitempty"`
}

// PetStatus pet status in the store
type PetStatus string

// Tag defines model for Tag.
type Tag struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// User defines model for User.
type User struct {
	Email      *string `json:"email,omitempty" validate:"email"`
	FirstName  *string `json:"firstName,omitempty"`
	Id         *int64  `json:"id,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	Password   *string `json:"password,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// FindPetsByStatusParams defines parameters for FindPetsByStatus.
type FindPetsByStatusParams struct {
	// Status Status values that need to be considered for filter
	Status *FindPetsByStatusParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// FindPetsByStatusParamsStatus defines parameters for FindPetsByStatus.
type FindPetsByStatusParamsStatus string

// FindPetsByTagsParams defines parameters for FindPetsByTags.
type FindPetsByTagsParams struct {
	// Tags Tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`
}

// DeletePetParams defines parameters for DeletePet.
type DeletePetParams struct {
	ApiKey *string `json:"api_key,omitempty"`
}

// UpdatePetWithFormParams defines parameters for UpdatePetWithForm.
type UpdatePetWithFormParams struct {
	// Name Name of pet that needs to be updated
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Status Status of pet that needs to be updated
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// UploadFileParams defines parameters for UploadFile.
type UploadFileParams struct {
	// AdditionalMetadata Additional Metadata
	AdditionalMetadata *string `form:"additionalMetadata,omitempty" json:"additionalMetadata,omitempty"`
}

// CreateUsersWithListInputJSONBody defines parameters for CreateUsersWithListInput.
type CreateUsersWithListInputJSONBody = []User

// LoginUserParams defines parameters for LoginUser.
type LoginUserParams struct {
	// Username The user name for login
	Username *string `form:"username,omitempty" json:"username,omitempty"`

	// Password The password for login in clear text
	Password *string `form:"password,omitempty" json:"password,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = Pet

// AddPetFormdataRequestBody defines body for AddPet for application/x-www-form-urlencoded ContentType.
type AddPetFormdataRequestBody = Pet

// UpdatePetJSONRequestBody defines body for UpdatePet for application/json ContentType.
type UpdatePetJSONRequestBody = Pet

// UpdatePetFormdataRequestBody defines body for UpdatePet for application/x-www-form-urlencoded ContentType.
type UpdatePetFormdataRequestBody = Pet

// PlaceOrderJSONRequestBody defines body for PlaceOrder for application/json ContentType.
type PlaceOrderJSONRequestBody = Order

// PlaceOrderFormdataRequestBody defines body for PlaceOrder for application/x-www-form-urlencoded ContentType.
type PlaceOrderFormdataRequestBody = Order

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// CreateUserFormdataRequestBody defines body for CreateUser for application/x-www-form-urlencoded ContentType.
type CreateUserFormdataRequestBody = User

// CreateUsersWithListInputJSONRequestBody defines body for CreateUsersWithListInput for application/json ContentType.
type CreateUsersWithListInputJSONRequestBody = CreateUsersWithListInputJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = User

// UpdateUserFormdataRequestBody defines body for UpdateUser for application/x-www-form-urlencoded ContentType.
type UpdateUserFormdataRequestBody = User
