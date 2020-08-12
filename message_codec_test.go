package gosmc

import (
	"reflect"
	"testing"

	"github.com/firmanmm/gosmc/encoder"
)

type MockValueEncoder struct {
}

func (m *MockValueEncoder) Encode(dataType encoder.ValueEncoderType, data interface{}) ([]byte, error) {
	return []byte{byte(data.(int))}, nil
}

func (m *MockValueEncoder) Decode(data []byte) (interface{}, error) {
	return int(data[0]), nil
}

func TestMessageCodecBehaviour(t *testing.T) {

	testData := []struct {
		Name     string
		Value    interface{}
		HasError bool
	}{
		{
			"Accepted",
			1,
			false,
		},
	}

	codec := NewSimpleMessageCodec()

	for _, val := range testData {
		t.Run(val.Name, func(t *testing.T) {
			encoded, err := codec.Encode(val.Value)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if reflect.DeepEqual(encoded, val.Value) {
				t.Errorf("Expected data to be transformed but nothing happens, %v", encoded)
			}
			decoded, err := codec.Decode(encoded)
			if err != nil != val.HasError {
				t.Errorf("Expected error value of %v but got %v", val.HasError, err != nil)
			}
			if !reflect.DeepEqual(val.Value, decoded) {
				t.Errorf("Expected %v but got %v", val.Value, decoded)
			}
		})
	}
}
