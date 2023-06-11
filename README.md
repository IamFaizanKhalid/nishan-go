<img width="200" src="https://d1sr9z1pdl3mb7.cloudfront.net/wp-content/uploads/2023/05/25124833/Nishan-Pakistan-1024x482.png" alt="Nishan Pakistan">

[![Tests](https://github.com/IamFaizanKhalid/nishan-go/actions/workflows/test.yml/badge.svg)](https://github.com/IamFaizanKhalid/nishan-go/actions/workflows/test.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/IamFaizanKhalid/nishan-go)](https://goreportcard.com/report/github.com/IamFaizanKhalid/nishan-go) [![Release](https://img.shields.io/github/v/release/IamFaizanKhalid/nishan-go.svg?style=flat-square)](https://github.com/IamFaizanKhalid/nishan-go/releases) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/IamFaizanKhalid/nishan-go/blob/master/LICENSE)
<hr>

Golang wrapper for [NADRA's Nishan](https://nishan.nadra.gov.pk/) APIs.

## Installation

```console
go get -u github.com/IamFaizanKhalid/nishan-go
```


## Usage Example

```go
package main

import (
	"fmt"
	"github.com/IamFaizanKhalid/nishan-go"
	"github.com/IamFaizanKhalid/nishan-go/errors"
)

func main() {
	api, err := nishan.NewClient("your_company", "your_api_key")
	if err != nil {
		panic(err)
	}

	response := api.KnowYourCustomer(nishan.Request{
		CitizenNo:   "1234567890124",
		Fingerprint: "/9j/4AAQSkZJRgABAQEAYABgAAD/2wBDAA...",
		FingerIndex: 6,
	})

	if response.ErrCode != errors.Nil {
		switch response.ErrCode {
		case errors.FingerprintMismatch:
			panic("fingerprint didn't match")
		default:
			panic(response.ErrMessage)
		}
	}

	fmt.Printf("Welcome %s %s!", response.PersonalData.Name, response.PersonalData.FatherName)
}
```


## Reference
- [API Docs](https://nishan.nadra.gov.pk/tech-stack)

## License

© 2023-time.Now() Faizan Khalid

Released under the [MIT License](https://github.com/IamFaizanKhalid/nishan-go/blob/master/LICENSE)
