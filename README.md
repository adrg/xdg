xdg
===
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/xdg)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/xdg)](https://goreportcard.com/report/github.com/adrg/xdg)

Provides an implementation of the [XDG Base Directory Specification](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).
The specification defines a set of standard paths for storing application files,
including data and configuration files. For portability and flexibility reasons,
applications should use the XDG defined locations instead of hardcoding paths.
The package also includes the locations of well known user directories, based
on [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).
The current implementation supports Linux, Windows and Mac OS.

Full documentation can be found at: https://godoc.org/github.com/adrg/xdg

## Installation
    go get github.com/adrg/xdg

## Usage

#### XDG Base Directory

```go
package main

import (
	"log"

	"github.com/adrg/xdg"
)

func main() {
	// XDG Base Directory paths.
	log.Println("Home config directory:", xdg.DataHome)
	log.Println("Data directories:", xdg.DataDirs)
	log.Println("Home config directory:", xdg.ConfigHome)
	log.Println("Config directories:", xdg.ConfigDirs)
	log.Println("Cache directory:", xdg.CacheHome)
	log.Println("Runtime directory:", xdg.RuntimeDir)

	// Obtain a suitable location for application config files.
	// ConfigFile takes one parameter which must contain the name of the file,
	// but it can also contain a set of parent directories. If the directories
	// don't exists, they will be created relative to the base config directory.
	configFilePath, err := xdg.ConfigFile("appname/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Save the config file at:", configFilePath)

	// For other types of application files use:
	// xdg.DataFile()
	// xdg.CacheFile()
	// xdg.RuntimeFile()

	// Finding application config files.
	// SearchConfigFile takes one parameter which must contain the name of
	// the file, but it can also contain a set of parent directories relative
	// to the config search paths (xdg.ConfigHome and xdg.ConfigDirs).
	configFilePath, err = xdg.SearchConfigFile("appname/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Config file was found at:", configFilePath)

	// For other types of application files use:
	// xdg.SearchDataFile()
	// xdg.SearchCacheFile()
	// xdg.SearchRuntimeFile()
}
```

#### XDG User Directories

```go
package main

import (
	"log"

	"github.com/adrg/xdg"
)

func main() {
	// XDG user directories.
	log.Println("Desktop directory:", xdg.UserDirs.Desktop)
	log.Println("Download directory:", xdg.UserDirs.Download)
	log.Println("Documents directory:", xdg.UserDirs.Documents)
	log.Println("Music directory:", xdg.UserDirs.Music)
	log.Println("Pictures directory:", xdg.UserDirs.Pictures)
	log.Println("Videos directory:", xdg.UserDirs.Videos)
	log.Println("Templates directory:", xdg.UserDirs.Templates)
	log.Println("Public directory:", xdg.UserDirs.PublicShare)

	// Obtain a suitable location for user document files.
	// DocumentFile takes one parameter which must contain the name of the file,
	// but it can also contain a set of parent directories. If the directories
	// don't exists, they will be created relative to the user's documents
	// directory.
	documentFilePath, err := xdg.UserDirs.DocumentFile("dir/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Save document file at:", documentFilePath)

	// For other types of user files use:
	// xdg.DesktopFile()
	// xdg.DownloadedFile()
	// xdg.AudioFile()
	// xdg.ImageFile()
	// xdg.VideoFile()
	// xdg.TemplateFile()
	// xdg.PublicFile()

	// Finding user document files.
	// SearchDocumentFile takes one parameter which must contain the name of
	// the file, but it can also contain a set of parent directories relative
	// to the user's document directory (xdg.UserDirs.Documents).
	documentFilePath, err = xdg.UserDirs.SearchDocumentFile("dir/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Document file was found at:", documentFilePath)

	// For other types of user files use:
	// xdg.SearchDesktopFile()
	// xdg.SearchDownloadedFile()
	// xdg.SearchAudioFile()
	// xdg.SearchImageFile()
	// xdg.SearchVideoFile()
	// xdg.SearchTemplateFile()
	// xdg.SearchPublicFile()
}
```

## References
For more information see the
[XDG Base Directory Specification](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html) and
[XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.
See [CONTRIBUTING.MD](https://github.com/adrg/xdg/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2014 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/xdg/blob/master/LICENSE) for more details.
