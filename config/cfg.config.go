package config

import (
	"encoding/json"
	"fmt"
	"go/tutorial/exception"
	"go/tutorial/models"
	"os"
	"strings"
)

var Cfg *models.Configuration

func InitCfg() {
	filePath := fmt.Sprintf("env/%s.json", appEnv())
	plan, readFileErr := os.ReadFile(filePath)
	exception.PanicLog(readFileErr)
	var config *models.Configuration
	err := json.Unmarshal(plan, &config)
	exception.PanicLog(err)
	Cfg = config
}
func appEnv() string {
	if os.Getenv("BUILD_STAGE") != "" {
		return strings.ToLower(os.Getenv("BUILD_STAGE"))
	} else if os.Getenv("GOLANG_ENVIRONMENT") != "" {
		return os.Getenv("GOLANG_ENVIRONMENT")
	}
	return "local"
}
