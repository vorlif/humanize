# Humanize ![Test status](https://github.com/vorlif/humanize/workflows/Test/badge.svg) [![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![PkgGoDev](https://pkg.go.dev/badge/github.com/vorlif/humanize)](https://pkg.go.dev/github.com/vorlif/humanize) [![Go Report Card](https://goreportcard.com/badge/github.com/vorlif/humanize)](https://goreportcard.com/report/github.com/vorlif/humanize) [![codecov](https://codecov.io/gh/vorlif/humanize/branch/main/graph/badge.svg?token=N1O0ZE1OFW)](https://codecov.io/gh/vorlif/humanize) ![MinVersion](https://img.shields.io/badge/Go-1.24+-blue)

The `humanize` package provides a collection of functions to convert Go data structures into a human-readable format.
It was widely adapted from the [Django project](https://github.com/django/django) and also uses the Django translations.

More than 90 languages are supported. You can find a list of all supported languages under [locale/](locale)

### Usage

To use the `humanize` package, you first need to load the languages you want to use.

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

### Examples

#### Relative time

```go
t := time.Now().Add(5 * time.Minute)
collection := humanize.MustNew(humanize.WithLocale(es.New(), zhHans.New()))

h := collection.CreateHumanizer(language.English)
fmt.Println(h.NaturalTime(t))
// Output: 5 minutes from now

h = collection.CreateHumanizer(language.SimplifiedChinese)
fmt.Println(h.NaturalTime(t))
// Output: 5分钟以后
```

#### Language-specific time output

```go
collection := humanize.MustNew(humanize.WithLocale(es.New(), be.New(), de.New()))

now := time.Now()

h := collection.CreateHumanizer(language.English)
fmt.Println(h.FormatTime(now, humanize.DateTimeFormat))
// May 15, 2022, 6 p.m.

h := collection.CreateHumanizer(language.MustParse("be"))
fmt.Println(h.FormatTime(now, humanize.DateTimeFormat))
// траўня 15, 2022, 6 папаўдні

h := collection.CreateHumanizer(language.German)
fmt.Println(h.FormatTime(now, humanize.DateTimeFormat))
// 15. Mai 2022 18:00
```

#### Intword

```go
collection := humanize.MustNew(humanize.WithLocale(es.New()))

h := collection.CreateHumanizer(language.English)
fmt.Println(h.Intword(1_000_000_000))
// Output: 1.0 billion

h = collection.CreateHumanizer(language.Spanish)
fmt.Println(h.Intword(1_000_000_000))
// Output: 1,0 millardo
```

#### Filesize

```go
collection := humanize.MustNew(humanize.WithLocale(de.New()))

h := collection.CreateHumanizer(language.English)
fmt.Println(h.FilesizeFormat(1024 * 1024 * 1024))
// Output: 1 GB
```

#### More examples

A collection of all functions and further examples can be found in
the [documentation](https://pkg.go.dev/github.com/vorlif/humanize).

### Migrate from `spreak`

To migrate from the spreak/humanize package, simply adjust the import path:

```diff
import (
-	"github.com/vorlif/spreak/humanize"
-	"github.com/vorlif/spreak/humanize/locale/ar"
+	"github.com/vorlif/humanize"
+	"github.com/vorlif/humanize/locale/ar"
)
```

### Add translations

If you would like to add a translation or add a new language, **do not do so in this repository**.
The translations in this repository are automatically generated from the Django translations and additions should also
be made there.
Use the following link to do so: https://www.transifex.com/django/django/.
For all non-translation related errors, this repository must be used.

## License

humanize is available under the MIT license. See the [LICENSE](LICENSE) file for more info.
The translations of the `locale` packages are licensed
under [Django's BSD license](https://raw.githubusercontent.com/django/django/main/LICENSE).