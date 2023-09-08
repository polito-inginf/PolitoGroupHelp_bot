package utils

import "fmt"

func LoadPortFromEnv(name string) string {
	return fmt.Sprintf(":%v", LoadEnv(name))
}
