package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFuncXHSHSTETTTETTE() {}

func TestFunctionName(t *testing.T) {
	expected := "github.com/DistilledP/postr/internal/test.testFuncXHSHSTETTTETTE"
	actual := GetFunctionName(testFuncXHSHSTETTTETTE)

	assert.Equal(t, expected, actual)
}
