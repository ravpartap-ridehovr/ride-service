package session

import (
	"context"
	json "encoding/json"
	"errors"
	"fmt"
	"time"

	hash "github.com/defensestation/goutils/hash"
	random "github.com/defensestation/goutils/random"
	redis "github.com/redis/go-redis/v9"
)

const (
	defaultVersion = 1
	defaultLimit   = 10
)

type Session struct {
	client *redis.Client
	Key    string
	Value  *Value
}

type Value struct {
	Version     int32
	AccessLimit int32
	Data        interface{}
	Metadata    map[string]string
}

type RideReqParams struct {
	IsRideReqReceived int32
	RideReqStatus     int
	RideReqUserid     string
}

func New(client *redis.Client) *Session {
	return &Session{client: client, Value: &Value{Metadata: map[string]string{}}}
}

func (s *Session) Set(ctx context.Context, exptime time.Duration) error {
	// set version if not provided
	if s.Value.Version == 0 {
		s.Value.Version = defaultVersion
	}

	if s.Value.AccessLimit == 0 {
		s.Value.AccessLimit = defaultLimit
	}

	// if token is not provided set random token
	if s.Key == "" {
		random, err := random.GenerateRandomString(21)
		if err != nil {
			return err
		}
		s.Key = random
	}

	// marshal token
	marshalValue, err := json.Marshal(&s.Value)
	if err != nil {
		return err
	}

	// set in redis
	if err := s.client.Set(ctx, s.Key, marshalValue, exptime).Err(); err != nil {
		return err
	}

	return nil
}

func (s *Session) Get(ctx context.Context, key string, skipupdate ...bool) error {
	s.Key = key
	// verify session against email
	fmt.Println(s.Key)
	response := s.client.Get(ctx, key)
	valueBytes, err := response.Bytes()
	if err != nil {
		return err
	}

	// unmarshal token
	if err := json.Unmarshal(valueBytes, s.Value); err != nil {
		return err
	}

	if len(skipupdate) == 0 {
		if err := s.Update(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (s *Session) Update(ctx context.Context) error {
	// if access limit is reached delete the key
	if s.Value.AccessLimit == 0 {
		// do not delete it just return error,
		s.client.Del(ctx, s.Key)
	}

	// if access limit is greater than 0 then decr
	if s.Value.AccessLimit > 0 {
		s.Value.AccessLimit -= 1
	}

	s.Value.Version += 1

	// marshal token
	marshalValue, err := json.Marshal(&s.Value)
	if err != nil {
		return err
	}

	// update session without resetting the ttl
	if err := s.client.SetXX(ctx, s.Key, marshalValue, -1).Err(); err != nil {
		return err
	}

	return nil
}

func CreateKey(ktype, email, accountid, username string) (string, error) {
	// set key.
	redisKey := hash.Hash256(email)
	if email == "" && accountid != "" && username != "" {
		redisKey = hash.Hash256(accountid + username)
	}

	// check if redis key is still empty return err.
	if redisKey == "" {
		return "", errors.New("key empty")
	}

	// append prefix.
	redisKey = ktype + "_" + redisKey

	return redisKey, nil
}
