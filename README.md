xdg
===

[![Build Status](https://github.com/adrg/xdg/workflows/CI/badge.svg)](https://github.com/adrg/xdg/actions?query=workflow%3ACI)
[![Code coverage](https://codecov.io/gh/adrg/xdg/branch/master/graphs/badge.svg?branch=master)](https://app.codecov.io/gh/adrg/xdg)
[![pkg.go.dev documentation](https://pkg.go.dev/badge/github.com/adrg/xdg)](https://pkg.go.dev/github.com/adrg/xdg)
[![MIT license](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go report card](https://goreportcard.com/badge/github.com/adrg/xdg)](https://goreportcard.com/report/github.com/adrg/xdg)
[![GitHub issues](https://img.shields.io/github/issues/adrg/xdg)](https://github.com/adrg/xdg/issues)
[![Buy me a coffee](https://img.shields.io/static/v1.svg?label=%20&message=Buy%20me%20a%20coffee&color=579fbf&logo=buy%20me%20a%20coffee&logoColor=white)](https://ko-fi.com/T6T72WATK)

Provides an implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html).
The specification defines a set of standard paths for storing application files,
including data and configuration files. For portability and flexibility reasons,
applications should use the XDG defined locations instead of hardcoding paths.

The package also includes the locations of well known [user directories](https://wiki.archlinux.org/index.php/XDG_user_directories)
and an implementation of the [state directory](https://wiki.debian.org/XDGBaseDirectorySpecification#Proposal:_STATE_directory) proposal.
Windows, macOS and most flavors of Unix are supported.

Full documentation can be found at: https://pkg.go.dev/github.com/adrg/xdg.

## Installation
    go get github.com/adrg/xdg

## Default locations

The package defines sensible defaults for XDG variables which are empty or not
present in the environment.

#### XDG Base Directory

|                 | Unix                                | macOS                           | Windows                                 |
| :-------------- | :---------------------------------- | :------------------------------ | :-------------------------------------- |
| XDG_DATA_HOME   | `~/.local/share`                    | `~/Library/Application Support` | `%LOCALAPPDATA%`                        |
| XDG_DATA_DIRS   | `/usr/local/share`<br/>`/usr/share` | `/Library/Application Support`  | `%APPDATA%\Roaming`<br/>`%PROGRAMDATA%` |
| XDG_CONFIG_HOME | `~/.config`                         | `~/Library/Preferences`         | `%LOCALAPPDATA%`                        |
| XDG_CONFIG_DIRS | `/etc/xdg`                          | `/Library/Preferences`          | `%PROGRAMDATA%`                         |
| XDG_CACHE_HOME  | `~/.cache`                          | `~/Library/Caches`              | `%LOCALAPPDATA%\cache`                  |
| XDG_RUNTIME_DIR | `/run/user/UID`                     | `~/Library/Application Support` | `%LOCALAPPDATA%`                        |

#### XDG user directories

|                     | Unix          | macOS         | Windows                   |
| :------------------ | :------------ | :------------ | :------------------------ |
| XDG_DESKTOP_DIR     | `~/Desktop`   | `~/Desktop`   | `%USERPROFILE%/Desktop`   |
| XDG_DOWNLOAD_DIR    | `~/Downloads` | `~/Downloads` | `%USERPROFILE%/Downloads` |
| XDG_DOCUMENTS_DIR   | `~/Documents` | `~/Documents` | `%USERPROFILE%/Documents` |
| XDG_MUSIC_DIR       | `~/Music`     | `~/Music`     | `%USERPROFILE%/Music`     |
| XDG_PICTURES_DIR    | `~/Pictures`  | `~/Pictures`  | `%USERPROFILE%/Pictures`  |
| XDG_VIDEOS_DIR      | `~/Videos`    | `~/Movies`    | `%USERPROFILE%/Videos`    |
| XDG_TEMPLATES_DIR   | `~/Templates` | `~/Templates` | `%USERPROFILE%/Templates` |
| XDG_PUBLICSHARE_DIR | `~/Public`    | `~/Public`    | `%PUBLIC%`                |

#### Non-standard directories

State directory

```
Unix:    ~/.local/state
macOS:   ~/Library/Application Support
Windows: %LOCALAPPDATA%
```

Application directories

```
Unix:
- $XDG_DATA_HOME/applications
- ~/.local/share/applications
- /usr/local/share/applications
- /usr/share/applications
- $XDG_DATA_DIRS/applications

macOS:
- /Applications

Windows:
- %APPDATA%\Roaming\Microsoft\Windows\Start Menu\Programs
```

Font directories

```
Unix:
- $XDG_DATA_HOME/fonts
- ~/.fonts
- ~/.local/share/fonts
- /usr/local/share/fonts
- /usr/share/fonts
- $XDG_DATA_DIRS/fonts

macOS:
- ~/Library/Fonts
- /Library/Fonts
- /System/Library/Fonts
- /Network/Library/Fonts

Windows:
- %windir%\Fonts
- %LOCALAPPDATA%\Microsoft\Windows\Fonts
```

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
	log.Println("Home data directory:", xdg.DataHome)
	log.Println("Data directories:", xdg.DataDirs)
	log.Println("Home config directory:", xdg.ConfigHome)
	log.Println("Config directories:", xdg.ConfigDirs)
	log.Println("Cache directory:", xdg.CacheHome)
	log.Println("Runtime directory:", xdg.RuntimeDir)

	// Non-standard directories.
	log.Println("Home state directory:", xdg.StateHome)
	log.Println("Application directories:", xdg.ApplicationDirs)
	log.Println("Font directories:", xdg.FontDirs)

	// Obtain a suitable location for application config files.
	// ConfigFile takes one parameter which must contain the name of the file,
	// but it can also contain a set of parent directories. If the directories
	// don't exist, they will be created relative to the base config directory.
	configFilePath, err := xdg.ConfigFile("appname/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Save the config file at:", configFilePath)

	// For other types of application files use:
	// xdg.DataFile()
	// xdg.CacheFile()
	// xdg.RuntimeFile()
	// xdg.StateFile()

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
	// xdg.SearchStateFile()
}
```

#### XDG user directories

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
}
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/adrg/xdg.svg)](https://starchart.cc/adrg/xdg)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome. See [CONTRIBUTING.MD](CONTRIBUTING.md).

**Contributors**:
[adrg](https://github.com/adrg),
[wichert](https://github.com/wichert),
[bouncepaw](https://github.com/bouncepaw),
[gabriel-vasile](https://github.com/gabriel-vasile),
[KalleDK](https://github.com/KalleDK).

## References

For more information see:
* [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html)
* [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories)
* [XDG state directory proposal](https://wiki.debian.org/XDGBaseDirectorySpecification#Proposal:_STATE_directory)
* [XDG_STATE_HOME proposal](https://lists.freedesktop.org/archives/xdg/2016-December/013803.html)

## License

Copyright (c) 2014 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.
