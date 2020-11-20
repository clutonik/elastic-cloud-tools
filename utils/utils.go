package utils

import (
	"errors"
	"fmt"
	"os"
)

// LookupEnvVars accepts a slice of environment variables as strings and returns a map with key as environment variable
// and value as its defined value. It also returns an error if any of the env vars is missing in environment.
func LookupEnvVars(vars []string) (map[string]string,error){
	// Create a map of env vars and their values
	envVarMap := make(map[string]string)

	// loop over all environment variables
	for _, v := range vars{
		value, ok := os.LookupEnv(v)
		if !ok {
			// Return error
			return envVarMap, errors.New(fmt.Sprintf("%v is not set, please set %v before executing this program", v, v))
		}
		envVarMap[v] = value // Add a key to map
	}
	// Return env var map
	return envVarMap,nil
}