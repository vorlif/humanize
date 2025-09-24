# Humanize ![Test status](https://github.com/vorlif/humanize/workflows/Test/badge.svg) [![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![PkgGoDev](https://pkg.go.dev/badge/github.com/vorlif/humanize)](https://pkg.go.dev/github.com/vorlif/humanize) [![Go Report Card](https://goreportcard.com/badge/github.com/vorlif/humanize)](https://goreportcard.com/report/github.com/vorlif/humanize) [![codecov](https://codecov.io/gh/vorlif/humanize/branch/main/graph/badge.svg?token=N1O0ZE1OFW)](https://codecov.io/gh/vorlif/humanize) ![MinVersion](https://img.shields.io/badge/Go-1.24+-blue)

The `humanize` package provides a collection of functions to convert Go data structures into a human-readable format.
It was widely adapted from the [Django project](https://github.com/django/django) and also uses the Django translations.
It should therefore be noted that the translations are under
the [Django's 3-clause BSD license](https://raw.githubusercontent.com/django/django/main/LICENSE).

### Usage

To use the `humanize` package, you first need to load the languages you want to use.
You can find a list of all supported languages under [locale/](locale)

```go
package main

import (
	"fmt"
	"time"

	"golang.org/x/text/language"

	"github.com/vorlif/humanize"
	"github.com/vorlif/humanize/locale/ar"
	"github.com/vorlif/humanize/locale/es"
	"github.com/vorlif/humanize/locale/zhHans"
)

func main() {
	// Load the translations for the desired languages
	collection := humanize.MustNew(
		humanize.WithLocale(es.New(), ar.New(), zhHans.New()),
	)

	// Create a humanizer.
	// A humanizer features a collection of humanize functions for a language.
	h := collection.CreateHumanizer(language.English)

	// Uses the functions...
	fmt.Println(h.Intword(1_000_000_000))
	// Output: 1.0 billion

	fmt.Println(h.NaturalDay(time.Now()))
	// Output: today

	t := time.Now().Add(5 * time.Minute)
	fmt.Println(h.NaturalTime(t))
	// Output: 5 minutes from now

	d := -80 * time.Hour
	fmt.Println(h.TimeSince(d))
	// Output: 3 days, 8 hours

	// ... for different languages
	h = collection.CreateHumanizer(language.Spanish)
	fmt.Println(h.TimeSince(d))
	// Output: 3 días, 8 horas
}
```

A collection of all functions and further examples can be found in
the [documentation](https://pkg.go.dev/github.com/vorlif/humanize).

### Add translations

If you would like to add a translation or add a new language, **do not do so in this repository**.
The translations in this repository are automatically generated from the Django translations and additions should also
be made there.
Use the following link to do so: https://www.transifex.com/django/django/.
For all non-translation related errors, this repository must be used.

## License

humanize is available under the MIT license. See the [LICENSE](LICENSE) file for more info.
The translations of the `humanize` packages are licensed
under [Django's BSD license](https://raw.githubusercontent.com/django/django/main/LICENSE).