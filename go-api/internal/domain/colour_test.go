package domain

import (
	"errors"
	"testing"
)

func TestColourFrom_ValidCodes(t *testing.T) {
	tests := []struct {
		code     string
		expected Colour
	}{
		{"#FFFFFF", ColourWhite},
		{"#FF5733", ColourRed},
		{"#FFC300", ColourOrange},
		{"#FFFF66", ColourYellow},
		{"#CCFF99", ColourGreen},
		{"#6666FF", ColourBlue},
		{"#9966CC", ColourPurple},
		{"#999999", ColourGrey},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			c, err := ColourFrom(tt.code)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if c != tt.expected {
				t.Errorf("got %v, want %v", c, tt.expected)
			}
		})
	}
}

func TestColourFrom_InvalidCode_ReturnsError(t *testing.T) {
	_, err := ColourFrom("#INVALID")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	var unsupported *UnsupportedColourError
	if !errors.As(err, &unsupported) {
		t.Fatalf("expected UnsupportedColourError, got %T", err)
	}
	if unsupported.Code != "#INVALID" {
		t.Errorf("expected code #INVALID, got %s", unsupported.Code)
	}
}

func TestColourFrom_EmptyCode_ReturnsDefault(t *testing.T) {
	c, err := ColourFrom("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Code != "#000000" {
		t.Errorf("expected #000000, got %s", c.Code)
	}
}

func TestColour_String(t *testing.T) {
	if ColourWhite.String() != "#FFFFFF" {
		t.Errorf("expected #FFFFFF, got %s", ColourWhite.String())
	}
}
