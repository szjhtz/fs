package test

import (
	"github.com/farseer-go/fs/parse"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {

	assert.Equal(t, 123, parse.ConvertValue("123", reflect.TypeOf(1)).Interface().(int))

	assert.Equal(t, 1, parse.Convert(1, 0))
	assert.Equal(t, int64(1), parse.Convert(1, int64(0)))
	assert.Equal(t, "1", parse.Convert(1, ""))
	assert.False(t, parse.Convert(0, false))
	assert.True(t, parse.Convert(int8(1), false))
	assert.True(t, parse.Convert(int16(1), false))
	assert.True(t, parse.Convert(int32(1), false))
	assert.True(t, parse.Convert(int64(1), false))
	assert.True(t, parse.Convert(uint(1), false))
	assert.True(t, parse.Convert(uint8(1), false))
	assert.True(t, parse.Convert(uint16(1), false))
	assert.True(t, parse.Convert(uint32(1), false))
	assert.True(t, parse.Convert(uint64(1), false))
	assert.True(t, parse.Convert(float32(1), false))
	assert.True(t, parse.Convert(float64(1), false))
	assert.False(t, parse.Convert(2, false))
	assert.True(t, parse.Convert(1, false))
	assert.Equal(t, int8(8), parse.Convert(8, int8(1)))
	assert.Equal(t, int16(8), parse.Convert(8, int16(1)))
	assert.Equal(t, int32(8), parse.Convert(8, int32(1)))
	assert.Equal(t, int64(8), parse.Convert(8, int64(1)))
	assert.Equal(t, uint(8), parse.Convert(8, uint(1)))
	assert.Equal(t, uint8(8), parse.Convert(8, uint8(1)))
	assert.Equal(t, uint16(8), parse.Convert(8, uint16(1)))
	assert.Equal(t, uint32(8), parse.Convert(8, uint32(1)))
	assert.Equal(t, uint64(8), parse.Convert(8, uint64(1)))
	assert.Equal(t, float32(8), parse.Convert(8, float32(1)))
	assert.Equal(t, float64(8), parse.Convert(8, float64(1)))
	assert.Equal(t, 8, parse.Convert(int8(8), 1))
	assert.Equal(t, 8, parse.Convert(int16(8), 1))
	assert.Equal(t, 8, parse.Convert(int32(8), 1))
	assert.Equal(t, 8, parse.Convert(int64(8), 1))
	assert.Equal(t, 8, parse.Convert(uint(8), 1))
	assert.Equal(t, 8, parse.Convert(uint8(8), 1))
	assert.Equal(t, 8, parse.Convert(uint16(8), 1))
	assert.Equal(t, 8, parse.Convert(uint32(8), 1))
	assert.Equal(t, 8, parse.Convert(uint64(8), 1))
	assert.Equal(t, 8, parse.Convert(float32(8), 1))
	assert.Equal(t, 8, parse.Convert(float64(8), 1))
	assert.Equal(t, "8", parse.Convert(int8(8), ""))
	assert.Equal(t, "8", parse.Convert(int16(8), ""))
	assert.Equal(t, "8", parse.Convert(int32(8), ""))
	assert.Equal(t, "8", parse.Convert(int64(8), ""))
	assert.Equal(t, "8", parse.Convert(uint(8), ""))
	assert.Equal(t, "8", parse.Convert(uint8(8), ""))
	assert.Equal(t, "8", parse.Convert(uint16(8), ""))
	assert.Equal(t, "8", parse.Convert(uint32(8), ""))
	assert.Equal(t, "8", parse.Convert(uint64(8), ""))
	assert.Equal(t, "8.1", parse.Convert(float32(8.1), ""))
	assert.Equal(t, "8.12", parse.Convert(float64(8.12), ""))

	assert.Equal(t, int8(8), parse.Convert("8", int8(1)))
	assert.Equal(t, int16(8), parse.Convert("8", int16(1)))
	assert.Equal(t, int32(8), parse.Convert("8", int32(1)))
	assert.Equal(t, int64(8), parse.Convert("8", int64(1)))
	assert.Equal(t, uint(8), parse.Convert("8", uint(1)))
	assert.Equal(t, uint8(8), parse.Convert("8", uint8(1)))
	assert.Equal(t, uint16(8), parse.Convert("8", uint16(1)))
	assert.Equal(t, uint32(8), parse.Convert("8", uint32(1)))
	assert.Equal(t, uint64(8), parse.Convert("8", uint64(1)))
	assert.Equal(t, float32(8), parse.Convert("8", float32(1)))
	assert.Equal(t, float64(8), parse.Convert("8", float64(1)))
	assert.True(t, parse.Convert("true", false))
	assert.True(t, parse.Convert("True", false))
	assert.False(t, parse.Convert("false", false))
	assert.False(t, parse.Convert("False", false))
	assert.Equal(t, "123", parse.Convert("123", ""))
	assert.Equal(t, 123, parse.Convert("123", 0))
	assert.Equal(t, 3, parse.Convert("123f", 3))
	assert.Equal(t, uint8(3), parse.Convert("123f", uint8(3)))
	assert.Equal(t, float32(3), parse.Convert("123f", float32(3)))
	assert.Equal(t, []int{1, 2, 3}, parse.Convert("1,2,3", []int{}))
	assert.Equal(t, []string{"1", "2", "3"}, parse.Convert("1,2,3", []string{}))

	assert.Equal(t, struct{}{}, parse.Convert("123", struct{}{}))

	assert.True(t, parse.Convert(true, false))
	assert.False(t, parse.Convert(false, false))
	assert.Equal(t, struct{}{}, parse.Convert(false, struct{}{}))
	assert.Equal(t, 1, parse.Convert(true, 0))
	assert.Equal(t, 0, parse.Convert(false, 0))
	assert.Equal(t, "true", parse.Convert(true, ""))
	assert.Equal(t, "false", parse.Convert(false, ""))
}
