package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemperatureTestSuite struct {
	suite.Suite
	Server *httptest.Server
	url    func(string) string
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
	suite.NoError(err)
	suite.Equal(http.StatusOK, resp.StatusCode)
}

func (suite *TemperatureTestSuite) TestCEPDoesNotExist() {
	notExistingCEP := "00000-000"
	resp, err := http.Get(suite.url(notExistingCEP))
	suite.NoError(err)
	suite.Equal(http.StatusNotFound, resp.StatusCode)
}

// // TestInvalidCEP tests the case where the CEP is invalid
func (suite *TemperatureTestSuite) TestNotValidCEP() {
	notValidCEP := "00003"
	resp, err := http.Get(suite.url(notValidCEP))
	suite.NoError(err)
	suite.Equal(http.StatusUnprocessableEntity, resp.StatusCode)
}

func TestTemperatureTestSuite(t *testing.T) {
	suite.Run(t, new(TemperatureTestSuite))
}
