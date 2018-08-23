package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestTagAddTag(t *testing.T) {
	tag := &QrPaymentTag{}
	tag.addTag("00", "Zero")
	assert.Equal(t, 1, len(tag.Tags))

	tag.addTag("01", "One")
	tag.addTag("02", "Two")
	assert.Equal(t, 3, len(tag.Tags))
}

func TestTagGetCode(t *testing.T) {
	tag := &QrPaymentTag{}
	assert.Equal(t, "", tag.getCode())

	tag.addTag("00", "01")
	assert.Equal(t, "000201", tag.getCode())

	tag.addTag("01", "12")
	assert.Equal(t, "000201010212", tag.getCode())

	tag30 := tag.addTag("30", "")
	tag30.addTag("00", "AID1234")
	tag30.addTag("01", "BID1234")
	tag30.addTag("02", "1234")
	tag30.addTag("03", "0")
	assert.Equal(t, "30350007AID12340107BID12340204123403010", tag30.getCode())
	assert.Equal(t, "00020101021230350007AID12340107BID12340204123403010", tag.getCode())
}

func TestTagGetTag(t *testing.T) {
	tag := &QrPaymentTag{}
	tag.addTag("00", "Zero")
	assert.Equal(t, "Zero", tag.getTag("00").Value)

	tag.addTag("01", "One")
	assert.Equal(t, "Zero", tag.getTag("00").Value)
	assert.Equal(t, "One", tag.getTag("01").Value)

	tag.addTag("02", "Two")
	assert.Equal(t, "Zero", tag.getTag("00").Value)
	assert.Equal(t, "One", tag.getTag("01").Value)
	assert.Equal(t, "Two", tag.getTag("02").Value)
}

func TestNewQrPaymentTag(t *testing.T) {
	tag := NewQrPaymentTag()
	code := tag.getCode()
	assert.Equal(t, "000201010212", code)

	qrCode := tag.getQrCode()
	assert.Equal(t, 20, len(qrCode))
	assert.True(t, strings.HasPrefix(qrCode, code))
}
