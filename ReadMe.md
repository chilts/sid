# sid : generate sortable identifiers

## Overview
[![GoDoc](https://godoc.org/github.com/chilts/sid?status.svg)](https://godoc.org/github.com/chilts/sid)
[![Build Status](https://travis-ci.org/chilts/sid.svg?branch=master)](https://travis-ci.org/chilts/sid)
[![Code Climate](https://codeclimate.com/github/chilts/sid/badges/gpa.svg)](https://codeclimate.com/github/chilts/sid)
[![Go Report Card](https://goreportcard.com/badge/github.com/chilts/sid)](https://goreportcard.com/report/github.com/chilts/sid)

This package is simple and only provides one function. The aim here is not pure speed, it is for an easy use-case
without having to worry about goroutines and locking.

## Install

```
go get github.com/chilts/sid
```

## Example

```
id1 := sid.Id()
id2 := sid.Id()

fmt.Printf("id1 = %s\n", id1)
fmt.Printf("id2 = %s\n", id2)

// -> "id1 = 1IeSBAWW83j-2wgJ4PUtlAr"
// -> "id2 = 1IeSBAWW9kK-0cDG64GQgGJ"
```

## Author

By [Andrew Chilton](https://chilts.org/), [@twitter](https://twitter.com/andychilton).

For [AppsAttic](https://appsattic.com/), [@AppsAttic](https://twitter.com/AppsAttic).

## License

[MIT](https://publish.li/mit-qLQqmVTO).

(Ends)
