package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Sample(name string) []byte {
	path, _ := filepath.Abs(fmt.Sprintf("../../../scripts/test/%s", name))
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Panicf("Cannot find file '%s' at '%s'", name, filepath.Dir(path))
	}

	return content
}
