package loader

import (
	"io"

	"github.com/duplicate-dirs-go/option"
)

type loader struct {
	opt *option.Option
}

func NewLoader(opt *option.Option) *loader {
	return &loader{opt}
}

func (l *loader) Load(w io.Writer, transformDataArr []string) error {
	for _, transformData := range transformDataArr {
		_, err := w.Write([]byte(transformData))
		if err != nil {
			return err
		}
	}
	return nil
}
