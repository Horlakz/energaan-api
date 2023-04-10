package app

import (
	"encoding/json"
	"log"

	validation "github.com/go-ozzo/ozzo-validation"
	appDto "github.com/horlakz/energaan-api/database/dto/app"
	"github.com/horlakz/energaan-api/validator"
)

type PlanValidator struct {
	validator.Validator[appDto.PlanDTO]
}

func (validator *PlanValidator) Validate(planDto appDto.PlanDTO) (map[string]interface{}, error) {
	err := validation.ValidateStruct(&planDto,
		validation.Field(&planDto.Slug,
			validation.Required,
			validation.Length(1, 32),
			validation.By(validator.ValidateDBUnique(planDto, "plans", "slug", map[string]interface{}{})),
		),
		validation.Field(&planDto.Title, validation.Required, validation.Length(1, 32)),
		validation.Field(&planDto.Image, validation.Required, validation.Length(1, 32)),
		validation.Field(&planDto.Description, validation.Required),
		validation.Field(&planDto.Price, validation.Required),
	)

	if err != nil {
		if e, ok := err.(validation.InternalError); ok {
			log.Println(e.InternalError())
			return nil, nil
		}

		var dat map[string]interface{}
		m, _ := json.Marshal(err)

		json.Unmarshal(m, &dat)
		return dat, err
	}

	return nil, nil
}

func (validator *PlanValidator) ValidateSlug(planDto appDto.PlanDTO) error {
	err := validation.Validate(
		&planDto.Slug,
		validation.Required,
		validation.Length(1, 32),
		validation.By(validator.ValidateDBUnique(planDto, "plans", "slug", map[string]interface{}{})),
	)

	if err != nil {
		if e, ok := err.(validation.InternalError); ok {
			log.Println(e.InternalError())
			return nil
		}

		return err
	}

	return nil
}
