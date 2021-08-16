# mailtrain-go-api

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=VBXHBYFU44T5W&source=url)
[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/virtomize/mailtrain-go-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/virtomize/mailtrain-go-api)](https://goreportcard.com/report/github.com/virtomize/mailtrain-go-api)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/virtomize/mailtrain-go-api/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/virtomize/mailtrain-go-api.svg?branch=master)](https://travis-ci.com/virtomize/mailtrain-go-api)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)

Implements the [mailtrain API](https://github.com/Mailtrain-org/mailtrain).

> !!! less tested !!!

## Installation

If you already installed GO on your system and configured it properly than its simply:

```
go get github.com/virtomize/mailtrain-go-api
```

If not follow [these instructions](https://golang.org/doc/install)

## Usage

### Simple example

```
  api, err := gomailtrain.NewAPI("https://mailtrain.example.com", "token")
  if err != nil {
    // handle error
  }

  // read all subscribed lists for an email
  lists, err := api.GetListsByEmail("mail@example.com")
  if err != nil {
    // handle error
  }
```

### Advanced examples

see [examples](https://github.com/virtomize/mailtrain-go-api/tree/master/examples) for some more usage examples

## Code Documentation

You find the full [code documentation here](https://godoc.org/github.com/virtomize/mailtrain-go-api).

## Contribution

Thank you for participating to this project.
Please see our [Contribution Guidlines](https://github.com/virtomize/mailtrain-go-api/blob/master/CONTRIBUTING.md) for more information.
