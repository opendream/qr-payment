package qr_payment

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
	q.AddTag("00", "01")
	q.AddTag("01", "12")
	return q
}

func (q *QrPaymentTag) AddTag(id string, value string) (*QrPaymentTag) {
	tag := &QrPaymentTag{
		Id: id,
		Value: value,
	}
	q.Tags = append(q.Tags, tag)
	return tag
}

func (q *QrPaymentTag) GetTag(id string) (*QrPaymentTag) {
	for _, tag := range q.Tags {
		if tag.Id == id {
			return tag
		}
	}
	return nil
}

func (q *QrPaymentTag) GetCode() string {
	if len(q.Tags) > 0 {
		var b bytes.Buffer
		for _, tag := range q.Tags {
			b.WriteString(tag.GetCode())
		}
		q.Value = b.String()
	}

	if q.Id == "" {
		return q.Value
	} else {
		return fmt.Sprintf(`%s%.2d%s`, q.Id, len(q.Value), q.Value)
	}
}

func (q *QrPaymentTag) GetQrCode() string {
	code := fmt.Sprintf(`%s6304`, q.GetCode())
	checksum := crc16.ChecksumCCITTFalse([]byte(code))
	qrCode := fmt.Sprintf(`%s%04X`, code, checksum)
	return qrCode
}
