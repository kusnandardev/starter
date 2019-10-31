package redisdb

// GetList :
func GetList(key string) ([]string, error) {
	list, err := rdb.SMembers(key).Result()
	return list, err
}

// RemoveList :
func RemoveList(key string, val interface{}) error {
	_, err := rdb.SRem(key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// AddList :
func AddList(key, val string) error {
	_, err := rdb.SAdd(key, val).Result()
	if err != nil {
		return err
	}
	return nil
}

// TurncateList :
func TurncateList(key string) error {
	_, err := rdb.Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}
