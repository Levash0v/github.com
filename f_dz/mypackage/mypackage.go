package mypackage

import (
	"testing"	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataAndProcess(t *testing.T) {
	mockFetcher := new(MockDataFetcher)
	mockFetcher.On("FetchData").Return("test data", nil)
	result, err := GetDataAndProcess(mockFetcher)
	assert.NoError(t, err)
	assert.Equal(t, "test data_processed", result)
	mockFetcher.AssertExpectations(t)
}