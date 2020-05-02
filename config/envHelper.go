package config

import (
	"os"
	"strconv"
)

var prefix string

func GetIntFromEnv(key string) (int, error){
	val, err := strconv.Atoi(os.Getenv(prefix + key))

	if err != nil {
		return 0, err
	}

	return val, nil
}

func GetBoolFromEnv(key string) bool{
	if os.Getenv(prefix+ key) == "true"{
		return true
	}

	return false
}

func GetStringFromEnv(key string) string{
	return os.Getenv(prefix + key)
}

func SetPrefix(key string){
	prefix = key + "_"
}


