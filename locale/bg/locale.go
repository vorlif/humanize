package bg

import (
	"embed"

	"golang.org/x/text/language"

	"github.com/vorlif/humanize"
)

//go:embed *.po
var fsys embed.FS

func New() *humanize.LocaleData {
	return &humanize.LocaleData{
		Lang: language.MustParse("bg"),
		Fs:   fsys,
		Format: &humanize.FormatData{
			DateFormat:      "d F Y",
			TimeFormat:      "H:i",
			MonthDayFormat:  "j F",
			ShortDateFormat: "d.m.Y",
		},
	}
}
