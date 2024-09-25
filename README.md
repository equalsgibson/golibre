<!-- markdownlint-configure-file { "MD004": { "style": "consistent" } } -->
<!-- markdownlint-disable MD033 -->

#

<p align="center">
  <picture>
    <img src="https://equalsgibson.github.io/golibre/golibre.png" width="512" height="512" alt="golibre logo">
  </picture>
    <br>
    <strong>Easily access Glucose Measurement data from the FreeStyle Libre systems</strong>

</p>

<!-- markdownlint-enable MD033 -->

-   **Easy to use**: Get up and running with the library in minutes

-   **Actively developed**: Ideas and contributions welcomed!

---

<div align="right">

[![Go][golang]][golang-url]
[![Code Coverage][coverage]][coverage-url]
[![Go Reference][goref]][goref-url]
[![Go Report Card][goreport]][goreport-url]

</div>

## Getting Started  

### Prerequisites  

Download and install Go, version 1.21+, from the [official Go website](https://go.dev/doc/install).

### Install  

```shell
go get github.com/equalsgibson/five9-go
```

#### Get the Glucose Data of a single patient

Below is a short example showing how to list all the users within your Five9 Domain using the library.

> **Note**  
> Make sure to `go get` the library, and set the required ENV variables (`LIBRELINKUP_EMAIL` and `LIBRELINKUP_PASSWORD`) before running the below example.

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/equalsgibson/golibre/golibre"
)

func main() {
	// Set up a new golibre service
	ctx := context.Background()
	service := golibre.NewService(
		"api.libreview.io",
		golibre.Authentication{
			Email:    os.Getenv("LIBRELINKUP_EMAIL"),    // Your email address
			Password: os.Getenv("LIBRELINKUP_PASSWORD"), // Your password
		},
	)

	connections, err := service.Connection().GetAllConnectionData(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print a count of all the patients that you are connected to, with a list of patient IDs
	log.Printf("You have %d patients that are sharing their data with you.\n\n", len(connections))

	for i, connection := range connections {
		log.Printf("\tPatient %d: %s\n", i, connection.PatientID)
	}
}
```

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, get inspired, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under a GNU License. See the `LICENSE` file for more information.

<!-- CONTACT -->

## Contact

[Chris Gibson (@equalsgibson)](https://github.com/equalsgibson)

Project Link: [https://github.com/equalsgibson/golibre](https://github.com/equalsgibson/golibre)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[golang]: https://img.shields.io/badge/v1.21-000?logo=go&logoColor=fff&labelColor=444&color=%2300ADD8
[golang-url]: https://go.dev/
[coverage]: https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fequalsgibson.github.io%2Ffive9-go%2Fcoverage%2Fcoverage.json&query=%24.total&label=Coverage
[coverage-url]: https://equalsgibson.github.io/five9-go/coverage/coverage.html
[goaction]: https://github.com/equalsgibson/five9-go/actions/workflows/go.yml/badge.svg?branch=main
[goaction-url]: https://github.com/equalsgibson/five9-go/actions/workflows/go.yml
[goref]: https://pkg.go.dev/badge/github.com/equalsgibson/five9-go.svg
[goref-url]: https://pkg.go.dev/github.com/equalsgibson/five9-go
[goreport]: https://goreportcard.com/badge/github.com/equalsgibson/five9-go
[goreport-url]: https://goreportcard.com/report/github.com/equalsgibson/five9-go
