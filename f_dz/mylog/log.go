package logs

import (
	"testing"	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func LogConfigsParams(prefix string, settings map[string]interface{}) {
	for key, value := range settings {
		switch value := value.(type) {
		case map[string]interface{}:
			if prefix == "" {
				LogConfigsParams(key, value)
			} else {
				LogConfigsParams(prefix+"."+key, value)
			}
		default:
			// Логирование простых значений
			if prefix == "" {
				Log.Infof("%s: %v", key, value)
			} else {
				Log.Infof("%s.%s: %v", prefix, key, value)
			}
		}
	}
}

func TestGetDataAndProcess(t *testing.T) {
	mockFetcher := new(MockDataFetcher)
	mockFetcher.On("FetchData").Return("test data", nil)
	result, err := GetDataAndProcess(mockFetcher)
	assert.NoError(t, err)
	assert.Equal(t, "test data_processed", result)
	mockFetcher.AssertExpectations(t)
}