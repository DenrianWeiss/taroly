package env

import (
	"os"
	"path/filepath"
)

var foundry = ""

func init() {
	env, b := os.LookupEnv("FOUNDRY_PATH_OVERRIDE")
	if b {
		foundry = env
	} else {
		foundry = ""
	}
}

func GetAnvilPath() string {
	return filepath.Join(foundry, "anvil")
}

func GetCastPath() string {
	return filepath.Join(foundry, "cast")
}
