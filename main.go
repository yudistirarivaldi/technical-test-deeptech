package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yudistirarivaldi/technical-test-deeptech/config"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/handler"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/middleware"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/repository"
	"github.com/yudistirarivaldi/technical-test-deeptech/internal/service"
)

type databaseConnections struct {
	mysql *sql.DB
}

type appServices struct {
	authService       *service.AuthService
	userService       *service.UserService
	categoriesService *service.CategoriesService
	productService    *service.ProductService
	transactionSerice *service.TransactionService
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbs, err := initDatabases(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize databases: %v", err)
	}
	defer closeDatabases(dbs)

	services := initServices(dbs, cfg)
	startHTTPServer(cfg, services)
}

func initDatabases(cfg *config.Config) (*databaseConnections, error) {
	dbs := &databaseConnections{}

	dbMysql, err := config.NewMySQLConnection(cfg.DatabaseMysql)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}
	dbs.mysql = dbMysql
	log.Println("Connected to MySQL")

	return dbs, nil
}

func closeDatabases(dbs *databaseConnections) {
	if dbs.mysql != nil {
		_ = dbs.mysql.Close()
	}
}

func initServices(dbs *databaseConnections, cfg *config.Config) *appServices {
	authRepo := repository.NewAuthRepository(dbs.mysql)
	userRepo := repository.NewUserRepository(dbs.mysql)
	categoriesRepo := repository.NewCategoriesRepository(dbs.mysql)
	productRepo := repository.NewProductRepository(dbs.mysql)
	transactionRepo := repository.NewTransactionRepository(dbs.mysql)

	authService := service.NewAuthService(authRepo, cfg.JWT.Secret)
	userService := service.NewUserService(userRepo)
	categoriesService := service.NewCategoriesService(categoriesRepo)
	productService := service.NewProductService(productRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	return &appServices{
		authService:       authService,
		categoriesService: categoriesService,
		productService:    productService,
		transactionSerice: transactionService,
		userService:       userService,
	}
}

func startHTTPServer(cfg *config.Config, services *appServices) {
	r := mux.NewRouter()

	authHandler := handler.NewAuthHandler(services.authService)
	categoriesHandler := handler.NewCategoriesHandler(services.categoriesService)
	productHandler := handler.NewProductHandler(services.productService)
	transactionHandler := handler.NewTransactionHandler(services.transactionSerice)
	userHandler := handler.NewUserHandler(services.userService)

	r.HandleFunc("/api/auth/register", authHandler.HandleRegister).Methods("POST")
	r.HandleFunc("/api/auth/login", authHandler.HandleLogin).Methods("POST")

	r.HandleFunc("/api/categories", middleware.JWTMiddleware(cfg.JWT.Secret, categoriesHandler.HandleInsert)).Methods("POST")
	r.HandleFunc("/api/categories", middleware.JWTMiddleware(cfg.JWT.Secret, categoriesHandler.HandleGetAll)).Methods("GET")
	r.HandleFunc("/api/categories/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, categoriesHandler.HandleGetByID)).Methods("GET")
	r.HandleFunc("/api/categories/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, categoriesHandler.HandleUpdate)).Methods("PUT")
	r.HandleFunc("/api/categories/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, categoriesHandler.HandleDelete)).Methods("DELETE")

	r.HandleFunc("/api/products", middleware.JWTMiddleware(cfg.JWT.Secret, productHandler.HandleInsert)).Methods("POST")
	r.HandleFunc("/api/products", middleware.JWTMiddleware(cfg.JWT.Secret, productHandler.HandleGetAll)).Methods("GET")
	r.HandleFunc("/api/products/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, productHandler.HandleGetByID)).Methods("GET")
	r.HandleFunc("/api/products/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, productHandler.HandleUpdate)).Methods("PUT")
	r.HandleFunc("/api/products/{id}", middleware.JWTMiddleware(cfg.JWT.Secret, productHandler.HandleDelete)).Methods("DELETE")

	r.Handle("/api/transactions", middleware.JWTMiddleware(cfg.JWT.Secret, transactionHandler.HandleCreate)).Methods("POST")
	r.Handle("/api/transactions/history", middleware.JWTMiddleware(cfg.JWT.Secret, transactionHandler.HandleGetUserTransactions)).Methods("GET")

	r.Handle("/api/users", middleware.JWTMiddleware(cfg.JWT.Secret, userHandler.HandleGetProfile)).Methods("GET")
	r.Handle("/api/users", middleware.JWTMiddleware(cfg.JWT.Secret, userHandler.HandleUpdateUser)).Methods("PUT")

	log.Printf("Server starting on port %s...", cfg.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Server.Port), r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
