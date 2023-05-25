package plan

import (
	"github.com/google/uuid"

	"github.com/horlakz/energaan-api/database/model"
)

type Gallery struct {
	model.Model
	Image       string    `json:"image"`
	Title       string    `json:"title"`
	CreatedByID uuid.UUID `gorm:"Column:created_by" json:"createdBy"`
	UpdatedByID uuid.UUID `gorm:"Column:updated_by" json:"updatedBy"`
	DeletedByID uuid.UUID `gorm:"Column:deleted_by" json:"deletedBy"`
}

func (Gallery) TableName() string {
	return "categories"
}
