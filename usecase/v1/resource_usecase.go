package v1

import (
	"github.com/faizalnurrozi/phincon-browse/domain/usecases"
	"github.com/faizalnurrozi/phincon-browse/domain/view_models"
	"github.com/faizalnurrozi/phincon-browse/packages/pokeapi-go"
	"github.com/faizalnurrozi/phincon-browse/usecase"
	"gitlab.com/s2-backend/packages/functioncaller"
	"gitlab.com/s2-backend/packages/logruslogger"
	"strings"
)

type ResourceUseCase struct {
	*usecase.Contract
}

func NewResourceUseCase(ucContract *usecase.Contract) usecases.IResourceUseCase {
	return &ResourceUseCase{Contract: ucContract}
}

// Browse all data by ordering and sorting
func (uc ResourceUseCase) Browse(page, limit int) (res []view_models.Resource, pagination view_models.PaginationVm, err error) {

	// Get page & limit
	offset, limit, page, _, _ := uc.SetPaginationParameter(page, limit, "orderBy", "sort")

	// Init data pokemon
	pokemons, err := pokeapi.Resource("pokemon", offset, limit)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "fetch-pokeapi-resource")
		return res, pagination, err
	}

	// Fetch data by resource
	for _, pokemon := range pokemons.Results {
		name := pokemon.Name
		url := pokemon.URL
		id := strings.Split(url, "/")[6]

		// Get picture
		pokemonPict, err := pokeapi.Pokemon(id)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "fetch-pokeapi-pokmeon")
			return res, pagination, err
		}

		res = append(res, view_models.Resource{
			ID:      id,
			Name:    name,
			Picture: pokemonPict.Sprites.FrontDefault,
		})
	}

	totalCount := pokemons.Count
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, err
}
