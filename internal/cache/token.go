package cache

import (
	"fmt"
	"github.com/consumer-superhero-register/internal/cache/model"
)

// SetToken stores Firebase Messaging Token into Redis cache.
func (c *Cache) SetToken(token model.FirebaseMessagingToken) error {
	if err := c.Redis.Set(fmt.Sprintf(c.TokenKeyFormat, token.SuperheroID), token, 0).Err(); err != nil {
		return err
	}

	return nil
}
