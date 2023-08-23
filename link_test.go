package link

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {

	html :=
		`<html>
		<body>
		  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
		</body>
		</html>`

	expected := []Link{{Href: "/dog-cat", Text: "dog cat"}}
	got, err := ParseHtml(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %#v, got %#v", expected, got)
	}

}
