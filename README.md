xdg
===
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/xdg)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/xdg)](https://goreportcard.com/report/github.com/adrg/xdg)

Provides an implementation of the XDG Base Directory Specification. The
specification defines a set of standard paths for storing application files
including data and configuration files. For portability and flexibility
reasons, applications should use the XDG defined locations instead of
hardcoding paths.

Full documentation can be found at: https://godoc.org/github.com/adrg/xdg

## Installation
    go get github.com/adrg/xdg

## Usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

func main() {
	// XDG directories
	fmt.Println(xdg.DataHome)
	fmt.Println(xdg.DataDirs)
	fmt.Println(xdg.ConfigHome)
	fmt.Println(xdg.ConfigDirs)
	fmt.Println(xdg.CacheHome)
	fmt.Println(xdg.RuntimeDir)

	// Finding a config file
	// The function takes one parameter which should contain the config
	// filename but it can also contain a set of directories relative to the
	// config search paths (xdg.ConfigHome and xdg.ConfigDirs).
	// Finding data, cache or runtime files is done in a similar manner using
	// xdg.SearchDataFile(), xdg.SearchCacheFile() and xdg.SearchRuntimeFile().
	configFilePath, err := xdg.SearchConfigFile("appname/config.yaml")
	if err != nil {
		// the config file was not found in the search paths
		// treat error
	} else {
		fmt.Println(configFilePath)
	}

	// Opening a config file
	// The function takes one parameter which should contain the config
	// filename but it can also contain a set of directories relative to the
	// config search paths. If the directories don't exists, they will be
	// created.
	// Opening data, cache or runtime files is done in a similar manner using
	// xdg.DataFile(), xdg.CacheFile() and xdg.RuntimeFile().
	configFilePath, err = xdg.ConfigFile("appname/config.yaml")
	if err != nil {
		// could not create path for the config file in any of the search paths
		// treat error
	} else {
		fp, err := os.Create(configFilePath)
		if err != nil {
			// could not open the file for some reason
			// treat error
		} else {
			// read/write the config file
		}
	}
}
```

## References
For more information see the
[XDG Base Directory Specification](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.
See [CONTRIBUTING.MD](https://github.com/adrg/xdg/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2014 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/xdg/blob/master/LICENSE) for more details.
