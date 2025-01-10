package redis

import (
	"context"
	"testing"
	"time"
)

func TestRedisTls(t *testing.T) {
	//tlsConfig := &tls.Config{
	//	InsecureSkipVerify: false, // Set to true if you want to skip certificate verification (not recommended for production)
	//}

	redisConfig := &Config{
		Addrs: []string{"localhost:6379"},
		//Password:  "yourpassword",
		DbIndex: 0,
		//TLSConfig: tlsConfig,
	}

	redisCmd, err := NewRedisCmd(redisConfig)
	if err != nil {
		panic(err)
	}
	_, err = redisCmd.Set(context.TODO(), "key", "hello", time.Minute*60).Result()
	if err != nil {
		panic(err)
	}
}
