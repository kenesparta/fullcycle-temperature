package integration

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemperatureTestSuite struct {
	suite.Suite
	url func(string) string
}

func (suite *TemperatureTestSuite) SetupTest() {
	suite.url = func(cep string) string {
		return fmt.Sprintf(
			"%s/%s?cep=%s",
			os.Getenv("API_URL"),
			"temperature",
			cep,
		)
	}
}

func (suite *TemperatureTestSuite) TestValidCEP() {
	validCEP := "79002-320"
	resp, err := http.Get(suite.url(validCEP))
	suite.T().Log("Executing: TestValidCEP")
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *TemperatureTestSuite) TestCEPDoesNotExist() {
	notExistingCEP := "00000-000"
	resp, err := http.Get(suite.url(notExistingCEP))
	suite.T().Log("Executing: TestCEPDoesNotExist")
	suite.NoError(err)
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}

func (suite *TemperatureTestSuite) TestNotValidCEP() {
	notValidCEP := "00003"
	resp, err := http.Get(suite.url(notValidCEP))
	suite.T().Log("Executing: TestNotValidCEP")
	suite.NoError(err)
	suite.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestTemperatureTestSuite(t *testing.T) {
	suite.Run(t, new(TemperatureTestSuite))
}
