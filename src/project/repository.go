package project

import (
	"go-just-portfolio/models"
)

type ProjectRepository interface {
	CreateTag(project_uuid string, tag string) (*string, error)
	CreateDescription(project_uuid, text string) (*string, error)
	DeleteprojectById(project_uuid, user_uuid string) error
	Newproject(user_uuid, category_uuid, title string) (*string, error)
	SetStateproject(state int, uuid, user_id string) error
	GetProjectsByShortname(shortname string) ([]models.Project, error)
	SavePhoto(project_uuid, photo_name, photo_type string) (*string, error)
	UpdatePrewiew(project_uuid, prewiew_name string) error
	RenameProject(user_uuid, uuid, title string) error
	GetProjectById(uuid string) models.InfoProjects

	AddDescription(project_uuid, text string) (*string, error)
	AddDescriptionIdToContent(project_uuid, description_uuid, type_ string) error
}
