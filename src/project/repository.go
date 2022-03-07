package project

import (
	"go-just-portfolio/models"
	"mime/multipart"
)

type ProjectRepository interface {
	CreateTag(project_uuid string, tag string) (*string, error)
	CreateDescription(project_uuid, key, value, lang string) (*string, error)
	DeleteprojectById(project_uuid, user_uuid string) error
	LoadPhoto(file *multipart.FileHeader) error
	Newproject(user_uuid, category_uuid, title string) (*string, error)
	SetStateproject(state int, uuid, user_id string) error
	GetProjectsByShortname(shortname string) ([]models.Project, error)
	SavePhoto(project_uuid, photo_name, photo_type string) error
}