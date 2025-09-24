package my

import (
	"embed"

	"golang.org/x/text/language"

	"github.com/vorlif/humanize"
)

//go:embed *.po
var fsys embed.FS

func New() *humanize.LocaleData {
	return &humanize.LocaleData{
		Lang: language.MustParse("my"),
		Fs:   fsys,
	}
}
