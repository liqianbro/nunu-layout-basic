package repository

import (
	"github.com/go-nunu/nunu-layout-basic/pkg/log"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger) (*Repository, func(), error) {
	db := newMySQL(logger)
	rdb := newRedis(logger)
	cleanup := func() {
		logger.Info("closing the data resources")

		// close db
		gdb, err := db.DB()
		if err != nil {
			logger.Error("close db error", zap.Error(err))
		}
		gdb.Close()
		// close redis
		rdb.Close()
	}
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}, cleanup, nil
}
func newMySQL(logger *log.Logger) *gorm.DB {
	// TODO: init db
	//db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//return db
	return &gorm.DB{}
}

func newRedis(logger *log.Logger) *redis.Client {
	// TODO: init redis
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     conf.GetString("data.redis.addr"),
	//	Password: conf.GetString("data.redis.password"),
	//	DB:       conf.GetInt("data.redis.db"),
	//})
	//return rdb
	return &redis.Client{}
}
