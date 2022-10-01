# Shimmering Bee: Callbacks

[![license](https://img.shields.io/github/license/shimmeringbee/callbacks.svg)](https://github.com/shimmeringbee/callbacks/blob/master/LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg)](https://github.com/RichardLitt/standard-readme)
[![Actions Status](https://github.com/shimmeringbee/callbacks/workflows/test/badge.svg)](https://github.com/shimmeringbee/callbacks/actions)

> Simple utility to provide callbacks against an event interface.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

Library that provides basic callback functionality for any number of different callback events, each called in order
of addition. No locking has been implemented, and assumes all callbacks will be added during initialisation. 

## Install

Add an import and most IDEs will `go get` automatically, if it doesn't `go build` will fetch.

```go
import "github.com/shimmeringbee/callbacks"
```

## Usage

**This libraries API is unstable and should not yet be relied upon.**

This piece of software is not ready for general use.

```
    type EventOne struct {}

    funcOne := func(ctx context.Context, event EventOne) error {
        // Do work
        return nil
    }

    cb := callbacks.Create()
    cb.Add(funcOne)

    err := cb.Call(context.Background(), EventOne{})
```

## Maintainers

[@pwood](https://github.com/pwood)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/shimmeringbee/callbacks/issues/new) or submit PRs.

All Shimmering Bee projects follow the [Contributor Covenant](https://shimmeringbee.io/docs/code_of_conduct/) Code of Conduct.

## License

   Copyright 2019-2020 Shimmering Bee Contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.