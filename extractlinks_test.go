package extractlinks

import (
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	htmlBody := strings.NewReader(`
		<html>
		<body>
		  <h1>Hello!</h1>
		  <a href="/another-page">
			A link to
			<span>internal page</span>
		  </a>
		  <a href="https://jsfuncc.com">A link to external page</a>
		</body>
		</html>
	`)

	links, err := All(htmlBody)

	if err != nil {
		t.Error("Err should be nil")
	}

	if len(links) != 2 {
		t.Error("Links count should be 2")
	}

	for i, link := range links {
		if i == 0 {
			if link.Href != "/another-page" {
				t.Error("Anchor link href is invalid")
			}
			if link.Text != "A link to internal page" {
				t.Error("Anchor link text is invalid")
			}
		}

		if i == 1 {
			if link.Href != "https://jsfuncc.com" {
				t.Error("Anchor link href is invalid")
			}
			if link.Text != "A link to external page" {
				t.Error("Anchor link text is invalid")
			}
		}
	}
}
