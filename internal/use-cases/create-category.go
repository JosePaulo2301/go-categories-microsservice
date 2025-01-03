package use_cases

import (
	"log"

	"github.com/jose.dev2301/go-categories-microsservice/internal/entities"
)

type crateCategoryUseCase struct {
	//db
}

func NewCreateCategoryUseCase() *crateCategoryUseCase {
	return &crateCategoryUseCase{}
}

func (u *crateCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: persist-entity-to-db
	log.Println(category)

	return nil

}
