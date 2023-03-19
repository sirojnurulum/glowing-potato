package http

import (
	"context"
	"flag"
	"fmt"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	orders3 "glowing-potato/cmd/http/handlers/orders"
	"glowing-potato/cmd/http/router"
	"glowing-potato/config"
	"glowing-potato/infra/context/handler"
	"glowing-potato/infra/context/repository"
	"glowing-potato/infra/context/service"
	"glowing-potato/infra/db"
	"glowing-potato/repositories/order_payment_methods"
	"glowing-potato/repositories/order_products"
	"glowing-potato/repositories/orders"
	orders2 "glowing-potato/services/orders"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	routerCMD = &cobra.Command{
		Use:   "serve-http",
		Short: "Run http server",
		Long:  "glowing-potato-api",
		RunE:  runHTTP,
	}
)

func runHTTP(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	cfg := config.InitConfig()

	potatodb, err := db.Open(&cfg.DB)
	if err != nil {
		fmt.Println(err)
	}

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*time.Duration(cfg.Server.GraceFulTimeout), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	//init repositories ctx
	repositoryCtx := initRepository(potatodb)
	//init services ctx
	serviceCtx := initService(repositoryCtx)
	//init handlers
	handlerCtx := initHandler(serviceCtx)
	//init router
	r := router.InitRouter(handlerCtx)

	// cors handlers
	corsHandler := cors.New(cors.Options{
		AllowedHeaders: []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedOrigins: []string{
			"http://localhost:9091",
			"http://localhost:8081",
		},
		OptionsPassthrough: false,
		AllowCredentials:   true,
	})

	//server conf
	srv := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      corsHandler.Handler(r),
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	log.Printf("API Listening on %s", cfg.Server.Addr)
	//run out the server in a goroutine so that it doesnt block
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//block untuk we receive our signal
	<-c

	//create deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return nil
}

func ServeHTTP() *cobra.Command {
	return routerCMD
}

func initRepository(db *db.DB) *repository.RepositoryCtx {
	ordersRepository := orders.NewOrdersRepository(db)
	orderProductRepository := order_products.NewOrderProductRepository(db)
	orderPaymentMethodsRepository := order_payment_methods.NewOrderPaymentMethodsRepository(db)
	return &repository.RepositoryCtx{
		DB:                            db,
		OrdersRepository:              ordersRepository,
		OrderProductRepository:        orderProductRepository,
		OrderPaymentMethodsRepository: orderPaymentMethodsRepository,
	}
}

func initService(repositoryCtx *repository.RepositoryCtx) *service.ServiceCtx {
	ordersService := orders2.NewOrdersService(repositoryCtx)
	return &service.ServiceCtx{
		OrdersService: ordersService,
	}
}

func initHandler(serviceCtx *service.ServiceCtx) *handler.HandlerCtx {
	ordersHandler := orders3.NewOrdersHandler(serviceCtx)
	return &handler.HandlerCtx{
		OrdersHandler: ordersHandler,
	}
}
