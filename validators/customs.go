package validators

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const TagMyCustom = "mycustom"

func RegisterCustomValidators(engine *gin.Engine) *gin.Engine {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		fmt.Println("validator not found, exit")
		os.Exit(2)
	}

	v.RegisterValidation(TagMyCustom, ValidatorMyCustom)

	return engine
}

func ValidatorMyCustom(fl validator.FieldLevel) bool {
	fmt.Println("MY CUSTOM - NOT CHECK")
	return false
}
