package v1

import (
	"github.com/faizalnurrozi/phincon-browse/domain/usecases"
	"github.com/faizalnurrozi/phincon-browse/domain/view_models"
	"github.com/faizalnurrozi/phincon-browse/packages/functioncaller"
	"github.com/faizalnurrozi/phincon-browse/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-browse/packages/pokeapi-go"
	"github.com/faizalnurrozi/phincon-browse/usecase"
	"strconv"
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

		pokemonID, _ := strconv.Atoi(id)

		res = append(res, view_models.Resource{
			ID:      pokemonID,
			Name:    name,
			Picture: pokemonPict.Sprites.FrontDefault,
		})
	}

	totalCount := pokemons.Count
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, err
}

// ReadBy pokemon detail by ID
func(uc ResourceUseCase) ReadBy(pokemonID string) (res view_models.ResourceDetail, err error){

	// Get data pokemon
	pokemon, err := pokeapi.Pokemon(pokemonID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "fetch-pokeapi-resource")
		return res, err
	}

	// Store moves data to struct
	var moves []string
	for _, pokemonMove := range pokemon.Moves {
		moves = append(moves, pokemonMove.Move.Name)
	}

	// Store types data to struct
	var types []string
	for _, pokemonType := range pokemon.Types {
		types = append(types, pokemonType.Type.Name)
	}

	res = view_models.ResourceDetail{
		ID:      pokemon.ID,
		Name:    pokemon.Name,
		Picture: pokemon.Sprites.FrontDefault,
		Moves:   moves,
		Types:   types,
	}

	return res, err
}
