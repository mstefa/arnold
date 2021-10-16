package bootstrap

import (
	"arnold/internal/platform/server"
	"context"
	"time"

	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	// mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	// db, err := sql.Open("mysql", mysqlURI)
	// if err != nil {
	// 	return err
	// }

	// var (
	// 	commandBus = inmemory.NewCommandBus()
	// 	eventBus   = inmemory.NewEventBus()
	// )

	// courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	// creatingCourseService := creating.NewCourseService(courseRepository)

	// createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"codely"`
	DbPass    string        `default:"codely"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"codely"`
	DbTimeout time.Duration `default:"5s"`
}
