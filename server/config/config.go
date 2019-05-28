package config

import "os"

var (
	// MainnetURL - URL for mainnet node
	MainnetURL = getEnv("MAINNET_URL", "https://mainnet.infura.io")
)

func getEnv(name, defaultValue string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
