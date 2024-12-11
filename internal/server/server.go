package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"os"
	"os/signal"
	"test/internal/cfg"
	userHandler "test/internal/http/user"
	"test/internal/middleware"
	userRepo "test/internal/repository/user"
	userServis "test/internal/service/user"
	"test/internal/txmanager"
	"test/pkg/err_handler"
)

type Server struct {
	cfg   *cfg.Config
	fiber *fiber.App
	db    *sqlx.DB
}

func NewServer(cfg *cfg.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg: cfg,
		fiber: fiber.New(
			fiber.Config{
				ErrorHandler: err_handler.ErrHandler,
			}),
		db: db,
	}
}

func (s *Server) Start() {
	trmanager := txmanager.NewTxManager(s.db)

	//repo
	userRepo := userRepo.New(trmanager)

	//service
	userServise := userServis.NewService(userRepo, s.cfg)

	//handler
	userHandlers := userHandler.NewHandler(userServise)

	//groups
	userGroup := s.fiber.Group("/user")

	mw := middleware.New(s.cfg)
	userHandler.MapUserRoutes(userGroup, userHandlers, mw)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)

	go func() {
		if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)); err != nil {
			panic(err)
		}
	}()

	<-exit

	if err := s.fiber.Shutdown(); err != nil {
		panic(err)
	}
}
