package utils

import (
	"os"
	"strings"
)

func GetEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func GetKVEnv(key string) map[string]string {
	val := os.Getenv(key)
	if val == "" {
		return map[string]string{}
	}
	kv := map[string]string{}
	for _, v := range strings.Split(val, ",") {
		kv[strings.SplitN(v, "=", 2)[0]] = strings.SplitN(v, "=", 2)[1]
	}
	return kv
}

func CastToIface(m map[string]string) map[string]interface{} {
	x := make(map[string]interface{})
	for k, v := range m {
		x[k] = v
	}
	return x
}
