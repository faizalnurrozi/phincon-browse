package domain

import (
	"database/sql"
	"github.com/faizalnurrozi/phincon-browse/packages/functioncaller"
	"github.com/faizalnurrozi/phincon-browse/packages/jwe"
	"github.com/faizalnurrozi/phincon-browse/packages/jwt"
	"github.com/faizalnurrozi/phincon-browse/packages/logruslogger"
	"github.com/faizalnurrozi/phincon-browse/packages/str"
	"github.com/faizalnurrozi/phincon-browse/usecase"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	svc_file_storage "gitlab.com/s2-backend/svc-file-storage/domain/protos"
	"google.golang.org/grpc"
	"os"
)

type Config struct {
	UcContract           *usecase.Contract
	DB                   *sql.DB
	SvcFileStorageClient svc_file_storage.FileStorageServiceClient
	JweCredential        jwe.Credential
	JwtCredential        jwt.JwtCredential
	JwtConfig            jwtFiber.Config
	Validator            *validator.Validate
	GrpClientConn        *grpc.ClientConn
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	//jwe credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	//jwt credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//jwt config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Claims:     &jwt.CustomClaims{},
	}

	res.Validator = ValidatorDriver

	return res, err
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	ValidatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(ValidatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(ValidatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		Translator = transEN
	case "id":
		Translator = transID
	}
}
