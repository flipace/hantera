package lib

import "io/ioutil"

// ReadFile : reads a file from a given path and returns its contents
func ReadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return dat
}
