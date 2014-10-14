package xdg_test

import (
	"fmt"

	"github.com/adrg/xdg"
)

func ExampleDataFile() {
	dataFilePath, err := xdg.DataFile("appname/appdata.data")
	if err != nil {
		// treat error
	}

	// the data file can be opened at the returned path
	fmt.Println("Data file location: ", dataFilePath)
}

func ExampleConfigFile() {
	configFilePath, err := xdg.ConfigFile("appname/config.yaml")
	if err != nil {
		// treat error
	}

	// the config file can be opened at the returned path
	fmt.Println("Config file location: ", configFilePath)
}

func ExampleCacheFile() {
	cacheFilePath, err := xdg.CacheFile("appname/app.cache")
	if err != nil {
		// treat error
	}

	// the cache file can be opened at the returned path
	fmt.Println("Cache file location: ", cacheFilePath)
}

func ExampleRuntimeFile() {
	runtimeFilePath, err := xdg.RuntimeFile("appname/app.pid")
	if err != nil {
		// treat error
	}

	// the runtime file can be opened at the returned path
	fmt.Println("Runtime file location: ", runtimeFilePath)
}

func ExampleSearchDataFile() {
	dataFilePath, err := xdg.SearchDataFile("appname/appdata.data")
	if err != nil {
		// the data file could not be found
		// treat error
	}

	// the data file was found at the returned path
	fmt.Println("The data file was found at: ", dataFilePath)
}

func ExampleSearchConfigFile() {
	configFilePath, err := xdg.SearchConfigFile("appname/config.yaml")
	if err != nil {
		// the config file could not be found
		// treat error
	}

	// the config file was found at the returned path
	fmt.Println("The config file was found at: ", configFilePath)
}

func ExampleSearchCacheFile() {
	cacheFilePath, err := xdg.SearchCacheFile("appname/app.cache")
	if err != nil {
		// the cache file could not be found
		// treat error
	}

	// the cache file was found at the returned path
	fmt.Println("The cache file was found at: ", cacheFilePath)
}

func ExampleSearchRuntimeFile() {
	runtimeFilePath, err := xdg.SearchRuntimeFile("appname/app.pid")
	if err != nil {
		// the runtime file could not be found
		// treat error
	}

	// the runtime file was found at the returned path
	fmt.Println("The runtime file was found at: ", runtimeFilePath)
}
