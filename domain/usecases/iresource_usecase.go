package usecases

import "github.com/faizalnurrozi/phincon-browse/domain/view_models"

type IResourceUseCase interface {
	Browse(page, limit int) (res []view_models.Resource, pagination view_models.PaginationVm, err error)

	ReadBy(pokemonID string) (res view_models.ResourceDetail, err error)
}
