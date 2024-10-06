package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/feliperdamaceno/go-orders-api/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Client *redis.Client
}

func orderIdToKey(id uint64) string {
	return fmt.Sprintf("order:%d", id)
}

var ErrNotExist = errors.New("order do not exist")

func (r *RedisRepository) Create(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIdToKey(order.Id)

	transaction := r.Client.TxPipeline()

	err = transaction.SetNX(ctx, key, string(data), 0).Err()
	if err != nil {
		transaction.Discard()
		return fmt.Errorf("failed to set order: %w", err)
	}

	err = transaction.SAdd(ctx, "orders", key).Err()
	if err != nil {
		transaction.Discard()
		return fmt.Errorf("failed to add order to set: %w", err)
	}

	if _, err := transaction.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute database transaction: %w", err)
	}

	return nil
}

type GetAllPage struct {
	Count  uint64
	Cursor uint64
}

type GetResult struct {
	Orders []model.Order
	Cursor uint64
}

func (r *RedisRepository) GetAll(ctx context.Context, page GetAllPage) (GetResult, error) {
	response := r.Client.SScan(ctx, "orders", page.Cursor, "*", int64(page.Count))

	keys, cursor, err := response.Result()
	if err != nil {
		return GetResult{}, fmt.Errorf("failed to get orders ids: %w", err)
	}

	if len(keys) == 0 {
		return GetResult{}, nil
	}

	result, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return GetResult{}, fmt.Errorf("failed to get orders: %w", err)
	}

	orders := make([]model.Order, len(result))

	for index, value := range result {
		value := value.(string)

		order := model.Order{}

		err := json.Unmarshal([]byte(value), &order)
		if err != nil {
			return GetResult{}, fmt.Errorf("failed to decode order: %w", err)
		}

		orders[index] = order
	}

	return GetResult{
		Orders: orders,
		Cursor: cursor,
	}, nil
}

func (r *RedisRepository) GetById(ctx context.Context, id uint64) (model.Order, error) {
	key := orderIdToKey(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.Order{}, ErrNotExist
	}

	if err != nil {
		return model.Order{}, fmt.Errorf("failed to get order: %w", err)
	}

	order := model.Order{}

	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to decode order: %w", err)
	}

	return order, nil
}

func (r *RedisRepository) UpdateById(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIdToKey(order.Id)

	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotExist
	}

	if err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}

func (r *RedisRepository) DeleteById(ctx context.Context, id uint64) error {
	key := orderIdToKey(id)

	transaction := r.Client.TxPipeline()

	err := transaction.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		transaction.Discard()
		return ErrNotExist
	}

	if err != nil {
		transaction.Discard()
		return fmt.Errorf("failed delete order: %w", err)
	}

	err = transaction.SRem(ctx, "orders", key).Err()
	if err != nil {
		transaction.Discard()
		return fmt.Errorf("failed to remove orders from set: %w", err)
	}

	if _, err := transaction.Exec(ctx); err != nil {
		return fmt.Errorf("failed to execute database transaction: %w", err)
	}

	return nil
}
