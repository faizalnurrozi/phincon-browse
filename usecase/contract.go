package usecase

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2-backend/packages/jwe"
	"gitlab.com/s2-backend/packages/jwt"
	"gitlab.com/s2-backend/packages/redis"
	"gitlab.com/s2-backend/packages/watermill"
	svc_file_storage "gitlab.com/s2-backend/svc-file-storage/domain/protos"
	"github.com/faizalnurrozi/phincon-browse/domain/view_models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Contract struct {
	ReqID                string
	UserID               string
	RoleID               string
	App                  *fiber.App
	DB                   *sql.DB
	Mongo                *mongo.Database
	TX                   *sql.Tx
	RedisClient          redis.RedisClient
	JweCredential        jwe.Credential
	JwtCredential        jwt.JwtCredential
	Validate             *validator.Validate
	Translator           ut.Translator
	SvcFileStorageClient svc_file_storage.FileStorageServiceClient
	Kafka                watermill.Kafka
}

const (
	defaultLimit    = 10
	maxLimit        = 50
	defaultOrderBy  = "id"
	defaultSort     = "asc"
	defaultLastPage = 0

	// Default for product
	DefaultDaysOfNewProduct = 30
)

func (uc Contract) SetPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) SetPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
