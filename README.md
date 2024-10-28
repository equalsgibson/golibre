<!-- markdownlint-configure-file { "MD004": { "style": "consistent" } } -->
<!-- markdownlint-disable MD033 -->

# golibre

<p align="center">
  <picture>
    <img src="https://equalsgibson.github.io/golibre/resources/golibre.png" width="512" height="512" alt="golibre logo">
  </picture>
    <br>
    <strong>Easily access Glucose Measurement data from the FreeStyle Libre systems</strong>

</p>

<!-- markdownlint-enable MD033 -->

- **Easy to use**: Get up and running with the library in minutes

- **Actively developed**: Ideas and contributions welcomed!

---

<div align="right">

[![Go][golang]][golang-url]
[![Code Coverage][coverage]][coverage-url]
[![Go Reference][goref]][goref-url]
[![Go Report Card][goreport]][goreport-url]

</div>

## Getting Started

### Prerequisites

- Download and install Go, version 1.22+, from the [official Go website](https://go.dev/doc/install).
- If you do not already have a LibreLinkUp account, create one by downloading the LibreLinkUp App from the [iOS App Store](https://apps.apple.com/us/app/librelinkup/id1234323923) or [Google Play Store](https://play.google.com/store/apps/details?id=org.nativescript.LibreLinkUp)

> [!TIP]
> To make sure that your account credentials will work with the library, you can download the [bruno](https://www.usebruno.com/) application or the [Postman](https://www.postman.com/) application and test the requests manually.
> If you choose to download and use bruno, see the ./ops/docs/bruno folder for example requests.

### Install

```shell
go get github.com/equalsgibson/golibre@v0.0.2-alpha
```

#### Get the Connections shared with your account

Below is a short example showing how to get the connections from your account

> [!NOTE]
> Make sure to `go get` the library, and set the required ENV variables (`EMAIL` and `PASSWORD`) before running the below example.

https://github.com/equalsgibson/golibre/blob/59176f70461221b2e021caf02e574bf3816a40de/examples/main.go#L1-L34

Expected Output:

```bash
cgibson@wsl-ubuntuNexus:~/git/libre/golibre$ go run examples/main.go
You have 1 patients that are sharing their data with you.
        -> Patient 1: ID: 12345678-1234-1234-abcd-0242ac110002
```

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, get inspired, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git add . && git commit -am 'Add some AmazingFeature'`)
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

[golang]: https://img.shields.io/badge/v1.22-000?logo=go&logoColor=fff&labelColor=444&color=%2300ADD8
[golang-url]: https://go.dev/
[coverage]: https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fequalsgibson.github.io%2Fgolibre%2Fcoverage%2Fcoverage.json&query=%24.total&label=Coverage
[coverage-url]: https://equalsgibson.github.io/golibre/coverage/coverage.html
[goref]: https://pkg.go.dev/badge/github.com/equalsgibson/golibre.svg
[goref-url]: https://pkg.go.dev/github.com/equalsgibson/golibre
[goreport]: https://goreportcard.com/badge/github.com/equalsgibson/golibre
[goreport-url]: https://goreportcard.com/report/github.com/equalsgibson/golibre
