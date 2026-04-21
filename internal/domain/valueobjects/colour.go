// Package valueobjects contains DDD value objects for the
// CleanArchitecture domain layer.
package valueobjects

import (
	domainerrors "github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/errors"
)

// Colour is a value object representing a supported hex colour code.
//
// Ported from src/Domain/ValueObjects/Colour.cs. Value-object equality
// is implemented via the comparable Code field: two Colour values are
// equal when their Code fields are equal (Go's built-in == operator).
type Colour struct {
	Code string
}

// Predefined supported colours.
//
// These mirror the static properties on the .NET Colour class.
var (
	ColourWhite  = Colour{Code: "#FFFFFF"}
	ColourRed    = Colour{Code: "#FF5733"}
	ColourOrange = Colour{Code: "#FFC300"}
	ColourYellow = Colour{Code: "#FFFF66"}
	ColourGreen  = Colour{Code: "#CCFF99"}
	ColourBlue   = Colour{Code: "#6666FF"}
	ColourPurple = Colour{Code: "#9966CC"}
	ColourGrey   = Colour{Code: "#999999"}
)

// supportedColours is the ordered list of all valid colour codes.
func supportedColours() []Colour {
	return []Colour{
		ColourWhite,
		ColourRed,
		ColourOrange,
		ColourYellow,
		ColourGreen,
		ColourBlue,
		ColourPurple,
		ColourGrey,
	}
}

// ColourFrom returns the Colour for the given hex code, or an
// UnsupportedColourError if the code is not in the supported set.
//
// Equivalent to Colour.From(string) in .NET.
func ColourFrom(code string) (Colour, error) {
	candidate := Colour{Code: code}
	for _, c := range supportedColours() {
		if c == candidate {
			return candidate, nil
		}
	}
	return Colour{}, domainerrors.NewUnsupportedColourError(code)
}

// String returns the hex code for this colour (implements fmt.Stringer).
// Equivalent to Colour.ToString() / the implicit string conversion.
func (c Colour) String() string {
	return c.Code
}
