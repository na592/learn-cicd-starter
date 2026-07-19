package auth

import (
"errors"
"net/http"
"strings"
"reflect"
"testing"
)


func TestGetAPIKey( headers http.header)
{
	got := GetAPIKey( headers )
	want := ([]string{''}, nil)

	if !reflect.DeepEqual(want, got) {
      t.Fatalf("expected: %v, got: %v", want, got)
   }
}
