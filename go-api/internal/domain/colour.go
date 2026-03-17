package domain

// Colour represents a colour value object with a hex code.
type Colour struct {
	Code string
}

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

var supportedColours = map[string]Colour{
	"#FFFFFF": ColourWhite,
	"#FF5733": ColourRed,
	"#FFC300": ColourOrange,
	"#FFFF66": ColourYellow,
	"#CCFF99": ColourGreen,
	"#6666FF": ColourBlue,
	"#9966CC": ColourPurple,
	"#999999": ColourGrey,
}

// ColourFrom creates a Colour from a hex code string.
// Returns an error if the code is not a supported colour.
// An empty code returns a default colour of #000000.
func ColourFrom(code string) (Colour, error) {
	if code == "" {
		return Colour{Code: "#000000"}, nil
	}
	c, ok := supportedColours[code]
	if !ok {
		return Colour{}, &UnsupportedColourError{Code: code}
	}
	return c, nil
}

// String returns the hex code of the colour.
func (c Colour) String() string { return c.Code }
