package gomailtrain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type apiTestValue struct {
	Input []string
	Error error
}

func TestNewAPI(t *testing.T) {
	testValues := []apiTestValue{
		{[]string{"", "token"}, fmt.Errorf("url or token not set")},
		{[]string{"test", ""}, fmt.Errorf("url or token not set")},
		{[]string{"https://test.test", "token"}, nil},
		{[]string{"test", "token"}, fmt.Errorf("parse \"test\": invalid URI for request")},
	}

	for _, test := range testValues {
		api, err := NewAPI(test.Input[0], test.Input[1])
		if err != nil {
			assert.Equal(t, test.Error.Error(), err.Error())
		} else {
			assert.Equal(t, test.Input[0], api.endPoint.String())
			assert.Equal(t, test.Input[1], api.token)
		}
	}
}

func TestSetDebug(t *testing.T) {
	assert.False(t, DebugFlag)
	SetDebug(true)
	assert.True(t, DebugFlag)
	SetDebug(false)
	assert.False(t, DebugFlag)
}
