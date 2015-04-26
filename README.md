GoBlocktrail
==========
Simple wrapper in Go for Blocktrail's API


## Installation


```bash
$ go get github.com/conejoninja/goblocktrail
```

## Documentation
See [Go Doc](http://godoc.org/github.com/conejoninja/goblocktrail) or [Go Walker](http://gowalker.org/github.com/conejoninja/goblocktrail) for usage and details.

## Example of use

```go
package main

import (
    "./goblocktrail"
    "fmt"
)

func main() {

    var api *goblocktrail.API

    api = goblocktrail.NewAPI("YOUR_API_KEY")

    res, err := api.Address("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T")
    fmt.Println("Address", res, err)

    res, err = api.Transactions("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T", 1, 200, "asc")
    fmt.Println("Transactions", res, err)

    res, err = api.UnconfirmedTransactions("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T", 1, 200, "asc")
    fmt.Println("UnconfirmedTransactions", res, err)

    res, err = api.UnspentOutputs("1NcXPMRaanz43b1kokpPuYDdk6GGDvxT2T", 1, 200, "asc")
    fmt.Println("UnspentOutputs", res, err)

    res, err = api.Block("290000")
    fmt.Println("Block", res, err)

    res, err = api.BlockByHeight(290000)
    fmt.Println("BlockByHeight", res, err)

    res, err = api.Block("0000000000000000fa0b2badd05db0178623ebf8dd081fe7eb874c26e27d0b3b")
    fmt.Println("Block", res, err)

    res, err = api.BlockTransactions("0000000000000000fa0b2badd05db0178623ebf8dd081fe7eb874c26e27d0b3b", 1, 2, "asc")
    fmt.Println("BlockTransactions", res, err)

    res, err = api.BlockTransactionsByHeight(290000, 1, 20, "asc")
    fmt.Println("BlockTransactionsByHeight", res, err)

    res, err = api.AllBlocks(1, 20, "asc")
    fmt.Println("AllBlocks", res, err)

    res, err = api.LatestBlock()
    fmt.Println("LatestBlock", res, err)

    res, err = api.Transaction("c326105f7fbfa4e8fe971569ef8858f47ee7e4aa5e8e7c458be8002be3d86aad")
    fmt.Println("Transaction", res, err)

    res, err = api.VerifyMessage("BlockTrail API verifying message", "1F26pNMrywyZJdr22jErtKcjF8R3Ttt55G", "H3y1vHBgBfV+7vTCnlhnA0dfacRsbucHxPWR/uFIB8yoQkfJ0oi71/FRze710sorujdC+LsvId7Jq7VV9g38tZw=")
    fmt.Println("VerifyMessage", res, err)


}
```

## Noted
I wouldn't use it for anything serious or important.

## Contributing to GoBlocktrail:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.

Feel free to donate some bits : 1Krm9w78fhEjxjE8Scnm3BqZyA9G527MQ8


## License

This is distributed under the Apache License v2.0

Copyright 2014 Daniel Esteban  -  conejo@conejo.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

