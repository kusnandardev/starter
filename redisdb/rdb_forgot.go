package redisdb

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// Forgot :
type Forgot struct {
	Email     string `json:"email"`
	UserName  string `json:"name"`
	ResetLink string `json:"reset_link"`
}

// StoreForgot :
func StoreForgot(data interface{}) error {
	var forgot Forgot

	err := mapstructure.Decode(data, &forgot)
	if err != nil {
		return err
	}

	mForgot := map[string]interface{}{
		"email_type": "forgot",
		"data":       forgot,
	}

	dForgot, err := json.Marshal(mForgot)
	if err != nil {
		return err
	}

	_, err = rdb.SAdd("starter_email", string(dForgot)).Result()
	if err != nil {
		return err
	}

	return nil
}
