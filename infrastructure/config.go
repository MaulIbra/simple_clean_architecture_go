package infrastructure

import "maulibra/enigma_school/api-master/utils"

type env struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	schemaName string
}

func Environment() *env {
	return &env{
		dbUser:     utils.GetEnv("--dbUser", "root"),
		dbPassword: utils.GetEnv("--dbPassword", ""),
		dbHost:     utils.GetEnv("--dbHost", "127.0.0.1"),
		dbPort:     utils.GetEnv("--dbPort", "3306"),
		schemaName: utils.GetEnv("--dbSchema", "enigma_school"),
	}
}
