package env

import (
	"github.com/joho/godotenv"
	"github.com/nanhao/highgin/err"
	"os"
)

func Load() *err.Err {
	err1 := godotenv.Load(".env")
	err2 := godotenv.Load(".env.default")
	if err1 == nil && err2 == nil {
		return &err.Err{
			Code: err.LACK_OF_ENV_FILE,
			Msg: "neight .env nor .env.default is found",
		}
	}
	return nil
}

func Get(key string) string {
	return os.Getenv(key)
}

func GetStrict(key string) (string, *err.Err) {
	val, found := os.LookupEnv(key)
	if !found {
		return "", &err.Err{
			Code: err.LACK_OF_ENV_KEY,
			Msg: "lack of " + key,
		}
	}
	return val, nil
}