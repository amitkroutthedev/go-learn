package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
*/

const url = "https://www.akrout.dev/"

func main() {
	fmt.Println("Welcome to web request")

	res, err := http.Get(url)
	checkNilError(err)

	fmt.Printf("Response is of type %T\n", res)
	fmt.Println("response:", res)

	defer res.Body.Close() //caller's responsibility to close the connection

	databytes, err := io.ReadAll(res.Body)

	checkNilError(err)
	content:=string(databytes)

	fmt.Println("Content:",content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
