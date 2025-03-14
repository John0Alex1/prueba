package User

import (
	"strconv"
	r "v1/src/common/response"
	usecase "v1/src/modules/products/usecase"

	"github.com/gofiber/fiber/v2"
)

type FindByIdProductController struct {
	Usecase *usecase.ProductsUsecase
	Result  *r.Result
}

func NewFindByIdProductController(usecase *usecase.ProductsUsecase, r *r.Result) *FindByIdProductController {
	return &FindByIdProductController{
		Usecase: usecase,
		Result:  r,
	}
}

func (ph *FindByIdProductController) ValidateRequest(c *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return 0, err
	}
	return id, err
}

func (ph *FindByIdProductController) Run(c *fiber.Ctx) (err error) {
	id, err := ph.ValidateRequest(c)
	if err != nil {
		return ph.Result.Error(c, err.Error())
	}
	product, err := ph.Usecase.FindById(id)
	if err != nil {
		return ph.Result.Error(c, err.Error())
	}

	ph.Result.Ok(c, product)
	return
}
