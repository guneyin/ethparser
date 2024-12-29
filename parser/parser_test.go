package parser

import (
	"github.com/guneyin/ethparser/testutils"
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
	testutils.ShouldNotBeEqual(t, cb, 0)
}

func TestTXParser_Subscribe(t *testing.T) {
	parser := New()

	sb := parser.Subscribe(testAddr)
	testutils.ShouldBeEqual(t, sb, true)

	// multiple subscribe request should not be fail
	sb = parser.Subscribe(testAddr)
	testutils.ShouldBeEqual(t, sb, true)
}

func TestTXParser_GetTransactions(t *testing.T) {
	parser := New()

	tsShouldFail := parser.GetTransactions(testAddr)
	testutils.ShouldBeNil(t, tsShouldFail)

	_ = parser.Subscribe(failTestAddr)
	setBlockNum(failTestBlockNum)
	tsShouldFail = parser.GetTransactions(failTestAddr)
	testutils.ShouldBeNil(t, tsShouldFail)

	_ = parser.Subscribe(testAddr)
	setBlockNum(testBlockNum)
	tsSuccess := parser.GetTransactions(testAddr)
	testutils.ShouldNotBeNil(t, tsSuccess)
}

func TestTXParser_getBlock(t *testing.T) {
	parser := New().(*TXParser)

	block, err := parser.getBlock(failTestBlockNum)
	testutils.ShouldBeNil(t, block)
	testutils.ShouldNotBeNil(t, err)

	block, err = parser.getBlock(testBlockNum)
	testutils.ShouldNotBeNil(t, block)
	testutils.ShouldBeNil(t, err)
}

func TestTXParser_utils(t *testing.T) {
	hexStr := "0x0444"
	num := hexToInt(hexStr)
	testutils.ShouldNotBeEqual(t, num, 0)

	hexStr = "invalid_hex"
	num = hexToInt(hexStr)
	testutils.ShouldBeEqual(t, num, 0)
}
