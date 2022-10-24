package utils

import (
	"testing"

	"github.com/aaronysj/RobotJarvis/pkg/utils"
)

func TestGetTokens(t *testing.T) {
	tokens := utils.GetTokens()
	if tokens == nil {
		t.Errorf("tokens is empty")
	}
}
