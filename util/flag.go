package util

import "io/ioutil"

// InputFile is a helper flag that automatically parses a
// file name into the InputFile variable
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
