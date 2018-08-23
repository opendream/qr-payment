package main

import (
	"fmt"
	"bytes"
	"github.com/howeyc/crc16"
)

type QrPaymentTag struct {
	Id string
	Value string
	Code string
	Tags []*QrPaymentTag
}

func NewQrPaymentTag() *QrPaymentTag {
	q := &QrPaymentTag{}
	q.addTag("00", "01")
	q.addTag("01", "12")
	return q
}

func (q *QrPaymentTag) addTag(id string, value string) (*QrPaymentTag) {
	tag := &QrPaymentTag{
		Id: id,
		Value: value,
	}
	q.Tags = append(q.Tags, tag)
	return tag
}

func (q *QrPaymentTag) getTag(id string) (*QrPaymentTag) {
	for _, tag := range q.Tags {
		if tag.Id == id {
			return tag
		}
	}
	return nil
}

func (q *QrPaymentTag) getCode() string {
	if len(q.Tags) > 0 {
		var b bytes.Buffer
		for _, tag := range q.Tags {
			b.WriteString(tag.getCode())
		}
		q.Value = b.String()
	}

	if q.Id == "" {
		return q.Value
	} else {
		return fmt.Sprintf(`%s%.2d%s`, q.Id, len(q.Value), q.Value)
	}
}

func (q *QrPaymentTag) getQrCode() string {
	code := q.getCode()
	checksum := crc16.ChecksumCCITTFalse([]byte(code))
	qrCode := fmt.Sprintf(`%s6304%04X`, code, checksum)
	return qrCode
}
