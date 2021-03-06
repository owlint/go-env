package env

import (
	"fmt"
	"os"
    "strconv"
)

func GetMandatoryEnv(envName string) string {
	env, present := os.LookupEnv(envName)

	if !present {
		panic(fmt.Sprintf("Env var %s is mandatory", envName))
	}

	return env
}

func GetDefaultEnv(envName, defaultValue string) string {
	env, present := os.LookupEnv(envName)

	if !present {
		return defaultValue
	}

	return env
}

func GetMandatoryIntFromEnv(envName string) int {
    value := GetMandatoryEnv(envName)

    intVal, err := strconv.Atoi(value)
    if err != nil {
        panic(err)
    }

	return intVal
}

func GetDefaultIntFromEnv(envName, defaultValue string) int {
    value := GetDefaultEnv(envName, defaultValue)

    intVal, err := strconv.Atoi(value)
    if err != nil {
        panic(err)
    }

	return intVal
}
