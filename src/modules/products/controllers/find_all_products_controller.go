package User

import (
	r "v1/src/common/response"
	usecase "v1/src/modules/products/usecase"

	"github.com/gofiber/fiber/v2"
)

type FindAllProductsController struct {
	Usecase *usecase.ProductsUsecase
	Result  *r.Result
}

func NewFindAllProductsController(usecase *usecase.ProductsUsecase, r *r.Result) *FindAllProductsController {
	return &FindAllProductsController{
		Usecase: usecase,
		Result:  r,
	}
}

func (ph *FindAllProductsController) Run(c *fiber.Ctx) (err error) {
	products, err := ph.Usecase.FindAll()
	if err != nil {
		return ph.Result.Error(c, err.Error())
	}

	ph.Result.Ok(c, products)
	return
}
