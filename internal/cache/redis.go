package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"forma/internal/domain"
	"forma/internal/models"

	"github.com/redis/go-redis/v9"
)

type LeadCache struct {
	rdb *redis.Client
}

func NewLeadCache(rdb *redis.Client) domain.LeadCache {
	return &LeadCache{rdb: rdb}
}

func allLeadsKey() string {
	return "leads:all"
}

func leadByIDKey(id uint) string {
	return fmt.Sprintf("leads:item:%d", id)
}

func (c *LeadCache) GetAll(ctx context.Context) ([]models.Lead, error) {
	val, err := c.rdb.Get(ctx, allLeadsKey()).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var leads []models.Lead
	if err := json.Unmarshal([]byte(val), &leads); err != nil {
		return nil, err
	}

	return leads, nil
}

func (c *LeadCache) SetAll(ctx context.Context, leads []models.Lead, ttl time.Duration) error {
	data, err := json.Marshal(leads)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, allLeadsKey(), data, ttl).Err()
}

func (c *LeadCache) DeleteAll(ctx context.Context) error {
	return c.rdb.Del(ctx, allLeadsKey()).Err()
}

func (c *LeadCache) GetByID(ctx context.Context, id uint) (*models.Lead, error) {
	val, err := c.rdb.Get(ctx, leadByIDKey(id)).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var lead models.Lead
	if err := json.Unmarshal([]byte(val), &lead); err != nil {
		return nil, err
	}

	return &lead, nil
}

func (c *LeadCache) SetByID(ctx context.Context, lead *models.Lead, ttl time.Duration) error {
	if lead == nil {
		return nil
	}

	data, err := json.Marshal(lead)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, leadByIDKey(lead.ID), data, ttl).Err()
}

func (c *LeadCache) DeleteByID(ctx context.Context, id uint) error {
	return c.rdb.Del(ctx, leadByIDKey(id)).Err()
}