package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetVarDefaultValue(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"

	if getVar(variableName, defaultValue) != defaultValue {
		t.Fail()
	}
}

func TestGetVarValueFromENV(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"
	originValue := "SUPPER+TEST"
	if err := os.Setenv(variableName, originValue); err != nil {
		log.Fatal(err)
	}

	if getVar(variableName, defaultValue) != originValue {
		t.Fail()
	}
}

func TestRouting_RateBTC(t *testing.T) {
	srv := httptest.NewServer(handlers())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/rate/btc", srv.URL))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "BitCoin to USD rate: 0.000000 $\n" {
		t.Fail()
	}
}
