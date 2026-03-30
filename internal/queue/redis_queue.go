package queue

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisQueue() *RedisQueue {
	return &RedisQueue{
		client: redis.NewClient(&redis.Options{
			Addr: "redis:6379",
		}),
		ctx: context.Background(),
	}
}

func (q *RedisQueue) Enqueue(queue string, job interface{}) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return q.client.LPush(q.ctx, queue, data).Err()
}

func (q *RedisQueue) DequeuePriority(queues []string) (string, error) {
	result, err := q.client.BRPop(q.ctx, 0, queues...).Result()
	if err != nil {
		return "", err
	}
	return result[1], nil
}
