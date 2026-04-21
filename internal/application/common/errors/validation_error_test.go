package errors_test

import (
	"sort"
	"testing"

	"github.com/danielcreevey-cell/sports-api-dotnet-demo/internal/application/common/errors"
)

func TestDefaultConstructorCreatesAnEmptyErrorDictionary(t *testing.T) {
	actual := errors.NewValidationError().Errors

	if len(actual) != 0 {
		t.Fatalf("expected empty errors map, got %d entries", len(actual))
	}
}

func TestSingleValidationFailureCreatesASingleElementErrorDictionary(t *testing.T) {
	failures := []errors.ValidationFailure{
		{PropertyName: "Age", ErrorMessage: "must be over 18"},
	}

	actual := errors.NewValidationErrorFromFailures(failures).Errors

	keys := mapKeys(actual)
	if !stringSlicesEqual(keys, []string{"Age"}) {
		t.Fatalf("expected keys [Age], got %v", keys)
	}
	if !stringSlicesEqual(actual["Age"], []string{"must be over 18"}) {
		t.Fatalf("expected Age messages [\"must be over 18\"], got %v", actual["Age"])
	}
}

func TestMultipleValidationFailureForMultiplePropertiesCreatesAMultipleElementErrorDictionaryEachWithMultipleValues(t *testing.T) {
	failures := []errors.ValidationFailure{
		{PropertyName: "Age", ErrorMessage: "must be 18 or older"},
		{PropertyName: "Age", ErrorMessage: "must be 25 or younger"},
		{PropertyName: "Password", ErrorMessage: "must contain at least 8 characters"},
		{PropertyName: "Password", ErrorMessage: "must contain a digit"},
		{PropertyName: "Password", ErrorMessage: "must contain upper case letter"},
		{PropertyName: "Password", ErrorMessage: "must contain lower case letter"},
	}

	actual := errors.NewValidationErrorFromFailures(failures).Errors

	keys := mapKeys(actual)
	if !stringSlicesEqualIgnoreOrder(keys, []string{"Password", "Age"}) {
		t.Fatalf("expected keys [Password, Age] (any order), got %v", keys)
	}

	if !stringSlicesEqualIgnoreOrder(actual["Age"], []string{
		"must be 25 or younger",
		"must be 18 or older",
	}) {
		t.Fatalf("unexpected Age messages: %v", actual["Age"])
	}

	if !stringSlicesEqualIgnoreOrder(actual["Password"], []string{
		"must contain lower case letter",
		"must contain upper case letter",
		"must contain at least 8 characters",
		"must contain a digit",
	}) {
		t.Fatalf("unexpected Password messages: %v", actual["Password"])
	}
}

func mapKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func stringSlicesEqualIgnoreOrder(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aa := append([]string(nil), a...)
	bb := append([]string(nil), b...)
	sort.Strings(aa)
	sort.Strings(bb)
	return stringSlicesEqual(aa, bb)
}
