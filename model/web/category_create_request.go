package web

type CategoryCreateRequest struct{
	Id int `json:"id"`
	Name string `validate:"required,min=3,max=10" json:"name"`
}