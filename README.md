# POC in Golang (Reserve with Google)

https://developers.google.com/maps-booking/guides/starter-integration/overview

This project is a demo reading and writing binary file using google protocol buffers.

### Notes
1. hello.go is the main script which read and write `feeds.bin` (binary wire format) from message like this:
```go
	object := Merchant{
		MerchantId: "merch1",
		Name:       "1234",
		Url:        "https://www.restaurantpublicwebsite.com",
	}
```
2. `feeds.proto` is language and platfrom neutral way of describing message structure. We use google protocol compiler to generate `feeds.pb.go` which is an interface that help us develop go app. 

Google Feeds Proto Bundle: (https://developers.google.com/maps-booking/reference/feed-specifications/feeds-proto-bundle)

It contains all necessary message structure used in Google Starter Integration. It has Merchant, Service, etc message sturcture.

This command doesn't generate executable program. It runs script with the dependent script.
``` sh
$ go run hello.go feeds.pb.go
```

- Golang
- Google Maps Booking API
- Google Protocol Buffers