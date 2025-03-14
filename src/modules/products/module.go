package products

import (
	types "v1/src/common/types"
	dao "v1/src/infraestructure/db/dao"

	controller "v1/src/modules/products/controllers"
	usecase "v1/src/modules/products/usecase"

	"net/http"

	"go.uber.org/fx"
)

func configureModuleRoutes(
	ctrlFindAllProducts *controller.FindAllProductsController,
	ctrlFindByIdProduct *controller.FindByIdProductController,
	h *types.HandlersStore,
) {
	handlersModuleUser := &types.SliceHandlers{
		Prefix: "products",
		Routes: []types.HandlerModule{
			{
				Route:   "",
				Method:  http.MethodGet,
				Handler: ctrlFindAllProducts.Run,
			},
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: ctrlFindByIdProduct.Run,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUser)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(dao.NewMySQLProductDao),
		fx.Provide(controller.NewFindAllProductsController),
		fx.Provide(controller.NewFindByIdProductController),
		fx.Provide(usecase.NewProductsUsecase),
		fx.Invoke(configureModuleRoutes),
	}
}
