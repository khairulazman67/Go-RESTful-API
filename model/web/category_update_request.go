package web

type CategoryUpdateRequest struct{
	Id int `validate:"required"`
	Name string	`validate:"required,min=3,max=10" json:"name"`
}