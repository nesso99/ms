# ms

A Golang version of [Tiny milisecond conversion utility](https://github.com/zeit/ms)

# Install

```bash
$ go get github.com/nesso99/ms
```

# Usage

```go
package main

import (
    "github.com/nesso99/ms"
)

func main() {
    ms.ParseMS("2 days")  // 172800000
    ms.ParseMS("1d")      // 86400000
    ms.ParseMS("10h")     // 36000000
    ms.ParseMS("2.5 hrs") // 9000000
    ms.ParseMS("2h")      // 7200000
    ms.ParseMS("1m")      // 60000
    ms.ParseMS("5s")      // 5000
    ms.ParseMS("1y")      // 31557600000
    ms.ParseMS("100")     // 100

    // convert in milliseconds to string in short type
    ms.ParseString(60000, false)        // "1m"
    ms.ParseString(2*60000, false)      // "2m"
    ms.ParseString(ms.ParseMS("10 hours"), false) // "10h"

    // convert in milliseconds to string in long type
    ms.ParseString(60000, true)         // "1 minute"
    ms.ParseString(2*60000, true)       // "2 minutes"
    ms.ParseString(ms.ParseMS("10 hours"), true)  // "10 hours"
}

```

# Test
```bash
$ go test
```
