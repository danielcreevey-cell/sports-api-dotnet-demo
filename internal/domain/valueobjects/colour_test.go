package valueobjects_test

import (
	"errors"
	"testing"

	domainerrors "github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/errors"
	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/domain/valueobjects"
)

// Ported from tests/Domain.UnitTests/ValueObjects/ColourTests.cs.

func TestColourFrom_ReturnsCorrectColourCode(t *testing.T) {
	const code = "#FFFFFF"

	colour, err := valueobjects.ColourFrom(code)
	if err != nil {
		t.Fatalf("ColourFrom(%q) returned unexpected error: %v", code, err)
	}
	if colour.Code != code {
		t.Errorf("Code = %q, want %q", colour.Code, code)
	}
}

func TestColour_StringReturnsCode(t *testing.T) {
	colour := valueobjects.ColourWhite

	if got := colour.String(); got != colour.Code {
		t.Errorf("String() = %q, want %q", got, colour.Code)
	}
}

func TestColourFrom_ReturnsUnsupportedColourErrorForInvalidCode(t *testing.T) {
	_, err := valueobjects.ColourFrom("##FF33CC")
	if err == nil {
		t.Fatal("ColourFrom with unsupported code returned nil error, want UnsupportedColourError")
	}

	var unsupported *domainerrors.UnsupportedColourError
	if !errors.As(err, &unsupported) {
		t.Fatalf("error type = %T, want *UnsupportedColourError", err)
	}
	if unsupported.Code != "##FF33CC" {
		t.Errorf("UnsupportedColourError.Code = %q, want %q", unsupported.Code, "##FF33CC")
	}
}

func TestColour_ComparableWithEqualityOperator(t *testing.T) {
	c1 := valueobjects.Colour{Code: "#FFFFFF"}
	c2 := valueobjects.Colour{Code: "#FFFFFF"}
	c3 := valueobjects.Colour{Code: "#AAAAAA"}

	if c1 != c2 {
		t.Errorf("expected c1 == c2 for equal codes")
	}
	if c1 == c3 {
		t.Errorf("expected c1 != c3 for different codes")
	}
}

func TestColourFrom_AcceptsAllSupportedCodes(t *testing.T) {
	supported := []valueobjects.Colour{
		valueobjects.ColourWhite,
		valueobjects.ColourRed,
		valueobjects.ColourOrange,
		valueobjects.ColourYellow,
		valueobjects.ColourGreen,
		valueobjects.ColourBlue,
		valueobjects.ColourPurple,
		valueobjects.ColourGrey,
	}
	for _, want := range supported {
		got, err := valueobjects.ColourFrom(want.Code)
		if err != nil {
			t.Errorf("ColourFrom(%q) returned error %v, want nil", want.Code, err)
			continue
		}
		if got != want {
			t.Errorf("ColourFrom(%q) = %v, want %v", want.Code, got, want)
		}
	}
}
