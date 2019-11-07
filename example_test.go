package xdg_test

import (
	"fmt"

	"github.com/adrg/xdg"
)

func ExampleDataFile() {
	dataFilePath, err := xdg.DataFile("appname/app.data")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save data file at:", dataFilePath)
}

func ExampleConfigFile() {
	configFilePath, err := xdg.ConfigFile("appname/app.yaml")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save config file at:", configFilePath)
}

func ExampleCacheFile() {
	cacheFilePath, err := xdg.CacheFile("appname/app.cache")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save cache file at:", cacheFilePath)
}

func ExampleRuntimeFile() {
	runtimeFilePath, err := xdg.RuntimeFile("appname/app.pid")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save runtime file at:", runtimeFilePath)
}

func ExampleSearchDataFile() {
	dataFilePath, err := xdg.SearchDataFile("appname/app.data")
	if err != nil {
		// The data file could not be found.
	}

	fmt.Println("The data file was found at:", dataFilePath)
}

func ExampleSearchConfigFile() {
	configFilePath, err := xdg.SearchConfigFile("appname/app.yaml")
	if err != nil {
		// The config file could not be found.
	}

	fmt.Println("The config file was found at:", configFilePath)
}

func ExampleSearchCacheFile() {
	cacheFilePath, err := xdg.SearchCacheFile("appname/app.cache")
	if err != nil {
		// The cache file could not be found.
	}

	fmt.Println("The cache file was found at:", cacheFilePath)
}

func ExampleSearchRuntimeFile() {
	runtimeFilePath, err := xdg.SearchRuntimeFile("appname/app.pid")
	if err != nil {
		// The runtime file could not be found.
	}

	fmt.Println("The runtime file was found at:", runtimeFilePath)
}

func ExampleUserDirectories_DesktopFile() {
	desktopFilePath, err := xdg.UserDirs.DesktopFile("dir/file.ext")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save desktop file at:", desktopFilePath)
}

func ExampleUserDirectories_DownloadedFile() {
	downloadedFilePath, err := xdg.UserDirs.DownloadedFile("dir/file.ext")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save downloaded file at:", downloadedFilePath)
}

func ExampleUserDirectories_DocumentFile() {
	documentFilePath, err := xdg.UserDirs.DocumentFile("dir/file.txt")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save document file at:", documentFilePath)
}

func ExampleUserDirectories_AudioFile() {
	audioFilePath, err := xdg.UserDirs.AudioFile("dir/file.mp3")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save audio file at:", audioFilePath)
}

func ExampleUserDirectories_ImageFile() {
	imageFilePath, err := xdg.UserDirs.ImageFile("dir/file.png")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save image file at:", imageFilePath)
}

func ExampleUserDirectories_VideoFile() {
	videoFilePath, err := xdg.UserDirs.VideoFile("dir/file.mkv")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save video file at:", videoFilePath)
}

func ExampleUserDirectories_TemplateFile() {
	templateFilePath, err := xdg.UserDirs.TemplateFile("dir/file.tpl")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save template file at:", templateFilePath)
}

func ExampleUserDirectories_PublicFile() {
	publicFilePath, err := xdg.UserDirs.PublicFile("dir/file.ext")
	if err != nil {
		// Treat error.
	}

	fmt.Println("Save public file at:", publicFilePath)
}

func ExampleUserDirectories_SearchDesktopFile() {
	desktopFilePath, err := xdg.UserDirs.SearchDesktopFile("dir/file.ext")
	if err != nil {
		// The desktop file could not be found.
	}

	fmt.Println("The desktop file was found at:", desktopFilePath)
}

func ExampleUserDirectories_SearchDownloadedFile() {
	downloadedFilePath, err := xdg.UserDirs.SearchDownloadedFile("dir/file.ext")
	if err != nil {
		// The downloaded file could not be found.
	}

	fmt.Println("The downloaded file was found at:", downloadedFilePath)
}

func ExampleUserDirectories_SearchDocumentFile() {
	documentFilePath, err := xdg.UserDirs.SearchDocumentFile("dir/file.txt")
	if err != nil {
		// The document file could not be found.
	}

	fmt.Println("The document file was found at:", documentFilePath)
}

func ExampleUserDirectories_SearchAudioFile() {
	audioFilePath, err := xdg.UserDirs.SearchAudioFile("dir/file.mp3")
	if err != nil {
		// The audio file could not be found.
	}

	fmt.Println("The audio file was found at:", audioFilePath)
}

func ExampleUserDirectories_SearchImageFile() {
	imageFilePath, err := xdg.UserDirs.SearchImageFile("dir/file.png")
	if err != nil {
		// The image file could not be found.
	}

	fmt.Println("The image file was found at:", imageFilePath)
}

func ExampleUserDirectories_SearchVideoFile() {
	videoFilePath, err := xdg.UserDirs.SearchVideoFile("dir/file.mkv")
	if err != nil {
		// The video file could not be found.
	}

	fmt.Println("The video file was found at:", videoFilePath)
}

func ExampleUserDirectories_SearchTemplateFile() {
	templateFilePath, err := xdg.UserDirs.SearchTemplateFile("dir/file.tpl")
	if err != nil {
		// The template file could not be found.
	}

	fmt.Println("The template file was found at:", templateFilePath)
}

func ExampleUserDirectories_SearchPublicFile() {
	publicFilePath, err := xdg.UserDirs.SearchPublicFile("dir/file.ext")
	if err != nil {
		// The public file could not be found.
	}

	fmt.Println("The public file was found at:", publicFilePath)
}
