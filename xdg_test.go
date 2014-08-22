package xdg

import "fmt"

func ExampleDataFile() {
	dataFilePath, err := DataFile("appname/appdata.data")
	if err != nil {
		// treat error
	}

	// the data file can be opened at the returned path
	fmt.Println("Data file location: ", dataFilePath)
}

func ExampleConfigFile() {
	configFilePath, err := DataFile("appname/config.yaml")
	if err != nil {
		// treat error
	}

	// the config file can be opened at the returned path
	fmt.Println("Config file location: ", configFilePath)
}

func ExampleCacheFile() {
	cacheFilePath, err := DataFile("appname/app.cache")
	if err != nil {
		// treat error
	}

	// the cache file can be opened at the returned path
	fmt.Println("Cache file location: ", cacheFilePath)
}

func ExampleRuntimeFile() {
	runtimeFilePath, err := DataFile("appname/app.pid")
	if err != nil {
		// treat error
	}

	// the runtime file can be opened at the returned path
	fmt.Println(runtimeFilePath)
}

func ExampleSearchDataFile() {
	dataFilePath, err := SearchDataFile("appname/appdata.data")
	if err != nil {
		// the data file could not be found
		// treat error
	}

	// the data file was found at the returned path
	fmt.Println(dataFilePath)
}

func ExampleSearchConfigFile() {
	configFilePath, err := SearchConfigFile("appname/config.yaml")
	if err != nil {
		// the config file could not be found
		// treat error
	}

	// the config file was found at the returned path
	fmt.Println(configFilePath)
}

func ExampleSearchCacheFile() {
	cacheFilePath, err := SearchCacheFile("appname/app.cache")
	if err != nil {
		// the cache file could not be found
		// treat error
	}

	// the cache file was found at the returned path
	fmt.Println(cacheFilePath)
}

func ExampleSearchRuntimeFile() {
	runtimeFilePath, err := SearchRuntimeFile("appname/app.pid")
	if err != nil {
		// the runtime file could not be found
		// treat error
	}

	// the runtime file was found at the returned path
	fmt.Println(runtimeFilePath)
}
