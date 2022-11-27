package router

import (
	"github.com/alofeoluwafemi/go-ethereum-api/app/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	api := app.Group("/api/v1", limiter.New())

	api.Get("/", controllers.RenderHello)

	deploy := api.Group("/deploy")
	deploy.Post("/usdc", controllers.DeployUSDC)
	deploy.Post("/factory", controllers.DeployFactory)

	api.Get("/wallet-logic-address", controllers.GetCustodianWalletLogicAddress)
	api.Get("/escrow-address", controllers.GetEscrowAddress)

	escrow := api.Group("/escrow")

	escrow.Post("/set-usdc-address", controllers.SetUSDCTokenAddress)
	api.Get("/get-usdc-address", controllers.GetUSDCAddress)

	api.Post("/factory/new-wallet/:uuid", controllers.NewWallet)
	api.Get("/factory/wallet/:uuid", controllers.GetWallet)

	api.Post("/wallet/order/new/:address", controllers.NewBuyOrder)
	api.Post("/wallet/balance/:address", controllers.GetWalletUSDCBalance)
}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}
