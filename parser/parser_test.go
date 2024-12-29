package parser

import (
	"reflect"
	"testing"
)

const (
	testAddr         = "0x51c72848c68a965f66fa7a88855f9f7784502a7f"
	failTestAddr     = "0x0"
	testBlockNum     = "0x1481c79"
	failTestBlockNum = "0x"
)

func setBlockNum(num string) {
	blockNum = num
}

func TestTXParser_GetCurrentBlock(t *testing.T) {
	parser := New()

	cb := parser.GetCurrentBlock()
	shouldNotBeEqual(t, cb, 0)
}

func TestTXParser_Subscribe(t *testing.T) {
	parser := New()

	sb := parser.Subscribe(testAddr)
	shouldBeEqual(t, sb, true)

	// multiple subscribe request should not be fail
	sb = parser.Subscribe(testAddr)
	shouldBeEqual(t, sb, true)
}

func TestTXParser_GetTransactions(t *testing.T) {
	parser := New()

	tsShouldFail := parser.GetTransactions(testAddr)
	shouldBeNil(t, tsShouldFail)

	_ = parser.Subscribe(failTestAddr)
	setBlockNum(failTestBlockNum)
	tsShouldFail = parser.GetTransactions(failTestAddr)
	shouldBeNil(t, tsShouldFail)

	_ = parser.Subscribe(testAddr)
	setBlockNum(testBlockNum)
	tsSuccess := parser.GetTransactions(testAddr)
	shouldNotBeNil(t, tsSuccess)
}

func TestTXParser_getBlock(t *testing.T) {
	parser := New().(*TXParser)

	block, err := parser.getBlock(failTestBlockNum)
	shouldBeNil(t, block)
	shouldNotBeNil(t, err)

	block, err = parser.getBlock(testBlockNum)
	shouldNotBeNil(t, block)
	shouldBeNil(t, err)
}

func TestTXParser_utils(t *testing.T) {
	hexStr := "0x0444"
	num := hexToInt(hexStr)
	shouldNotBeEqual(t, num, 0)

	hexStr = "invalid_hex"
	num = hexToInt(hexStr)
	shouldBeEqual(t, num, 0)
}

func shouldBeEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func shouldNotBeEqual(t *testing.T, actual, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
	}
}

func shouldNotBeNil(t *testing.T, v any) {
	if isNil(v) {
		t.Errorf("should be nil, but got: %#v", v)
	}
}

func shouldBeNil(t *testing.T, v any) {
	if !isNil(v) {
		t.Error("should not be nil")
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	switch value.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:

		return value.IsNil()
	default:
		return false
	}
}
