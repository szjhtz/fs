package parse

import (
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("int转bool", func(t *testing.T) {
		result := Convert(0, false)
		if result != false {
			t.Error()
		}
		result = Convert(1, false)
		if result != true {
			t.Error()
		}
	})

	t.Run("int转字符串", func(t *testing.T) {
		result := Convert(1, "")
		if result != "1" {
			t.Error()
		}
	})

	t.Run("int转int64", func(t *testing.T) {
		result := Convert(1, int64(0))
		if result != int64(1) {
			t.Error()
		}
	})

	t.Run("字符串转bool", func(t *testing.T) {
		result := Convert("true", false)
		if result != true {
			t.Error()
		}
	})

	t.Run("字符串转int", func(t *testing.T) {
		result := Convert("123", 0)
		if result != 123 {
			t.Error()
		}
	})

	t.Run("字符串转int64", func(t *testing.T) {
		result := Convert("123", int64(0))
		if result != int64(123) {
			t.Error()
		}
	})
}
