package qr_payment

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestTagAddTag(t *testing.T) {
	tag := &QrPaymentTag{}
	tag.AddTag("00", "Zero")
	assert.Equal(t, 1, len(tag.Tags))

	tag.AddTag("01", "One")
	tag.AddTag("02", "Two")
	assert.Equal(t, 3, len(tag.Tags))
}

func TestTagGetCode(t *testing.T) {
	tag := &QrPaymentTag{}
	assert.Equal(t, "", tag.GetCode())

	tag.AddTag("00", "01")
	assert.Equal(t, "000201", tag.GetCode())

	tag.AddTag("01", "12")
	assert.Equal(t, "000201010212", tag.GetCode())

	tag30 := tag.AddTag("30", "")
	tag30.AddTag("00", "AID1234")
	tag30.AddTag("01", "BID1234")
	tag30.AddTag("02", "1234")
	tag30.AddTag("03", "0")
	assert.Equal(t, "30350007AID12340107BID12340204123403010", tag30.GetCode())
	assert.Equal(t, "00020101021230350007AID12340107BID12340204123403010", tag.GetCode())
}

func TestTagGetTag(t *testing.T) {
	tag := &QrPaymentTag{}
	tag.AddTag("00", "Zero")
	assert.Equal(t, "Zero", tag.GetTag("00").Value)

	tag.AddTag("01", "One")
	assert.Equal(t, "Zero", tag.GetTag("00").Value)
	assert.Equal(t, "One", tag.GetTag("01").Value)

	tag.AddTag("02", "Two")
	assert.Equal(t, "Zero", tag.GetTag("00").Value)
	assert.Equal(t, "One", tag.GetTag("01").Value)
	assert.Equal(t, "Two", tag.GetTag("02").Value)
}

func TestNewQrPaymentTag(t *testing.T) {
	tag := NewQrPaymentTag()
	code := tag.GetCode()
	assert.Equal(t, "000201010212", code)

	qrCode := tag.GetQrCode()
	assert.Equal(t, 20, len(qrCode))
	assert.True(t, strings.HasPrefix(qrCode, code))
}
