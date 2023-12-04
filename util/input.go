package util

import (
	"flag"
	"io/ioutil"
)

func HandleInput() string {
	var input InputFile
	flag.Var(&input, "input", "input file")
	flag.Parse()
	return input.String()
}

// InputFile is a helper flag that automatically parses a
// file flag [-inputFile value] into the InputFile variable
type InputFile string

func (i *InputFile) String() string {
	return string(*i)
}

func (t *InputFile) Set(f string) error {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	*t = InputFile(b)
	return nil
}
