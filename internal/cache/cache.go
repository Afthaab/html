package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/afthaab/job-portal/internal/models"
	"github.com/redis/go-redis/v9"
)

type RDBLayer struct {
	rdb *redis.Client
}

//go:generate mockgen -source=cache.go -destination=mockmodels/cache_mock.go -package=mockmodels

type Caching interface {
	AddToTheCache(ctx context.Context, jid uint, jobData models.Jobs) error
	GetTheCacheData(ctx context.Context, jid uint) (string, error)
	CacheTheOtp(ctx context.Context, key string, otp string) error
	CheckTheOtp(ctx context.Context, key string) (string, error)
	DeleteTheCache(ctx context.Context, key string) error
}

func NewRDBLayer(rdb *redis.Client) Caching {
	return &RDBLayer{
		rdb: rdb,
	}
}

func (c *RDBLayer) DeleteTheCache(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()
	return err
}

func (c *RDBLayer) CheckTheOtp(ctx context.Context, key string) (string, error) {
	str, err := c.rdb.Get(ctx, key).Result()
	return str, err
}

func (c *RDBLayer) CacheTheOtp(ctx context.Context, key string, otp string) error {
	// val, err := json.Marshal(otp)
	// if err != nil {
	// 	return err
	// }
	err := c.rdb.Set(ctx, key, otp, 1*time.Minute).Err()
	return err
}

func (c *RDBLayer) AddToTheCache(ctx context.Context, jid uint, jobData models.Jobs) error {
	jobID := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobData)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, jobID, val, 1*time.Minute).Err()
	return err
}

func (c *RDBLayer) GetTheCacheData(ctx context.Context, jid uint) (string, error) {
	jobID := strconv.FormatUint(uint64(jid), 10)
	str, err := c.rdb.Get(ctx, jobID).Result()
	return str, err
}
