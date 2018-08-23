# QR Payment

## Installation

`$ go get -u github.com/opendream/qr-payment`

## Example for Thai QR Payment

### Import package

`import "github.com/opendream/qr-payment`

### Create QR Tag

```
qr := qr_payment.NewQrPaymentTag()
tag30 := qr.AddTag("30", "")
tag30.AddTag("00", "AID")
tag30.AddTag("01", "BillerID")
tag30.AddTag("02", "Ref1")
tag30.AddTag("03", "Ref2")
qr.AddTag("53", "764")
qr.AddTag("54", "1200.00")
qr.AddTag("58", "TH")
qr.AddTag("59", "Opendream")
// Addtional tag.
tag62 := qr.AddTag("62", "")
tag62.AddTag("07", "DOD00000001")

code := qr.GetQrCode()
```

### Generate QR Image

```
import (
  "github.com/skip2/go-qrcode"
)

...

code := qr.GetQrCode()
png, err := qrcode.Encode(code, qrcode.Medium, 256)
if err != nil {
  // Do something
  return
}
```

### Generate QR Image Data URI

```
dataUri := fmt.Sprintf(`data:image/png;base64,%s`, base64.StdEncoding.EncodeToString(png))
```
