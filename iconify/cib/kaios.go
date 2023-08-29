package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kaios(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.958 0a3.745 3.745 0 0 0-3.786 3.708v24.573A3.732 3.732 0 0 0 5.958 32a3.737 3.737 0 0 0 3.802-3.708V3.709A3.776 3.776 0 0 0 5.958.001zm12.537 11.385a3.735 3.735 0 0 0-5.188-.88a3.577 3.577 0 0 0-.906 5.068l10.667 14.865a3.741 3.741 0 0 0 5.172.88a3.574 3.574 0 0 0 .99-4.958l-.083-.12l-10.641-14.854zm4.864-1.572a4.95 4.95 0 0 0 4.99-4.891a4.95 4.95 0 0 0-4.99-4.896a4.95 4.95 0 0 0-5 4.896a4.949 4.949 0 0 0 5 4.891z"/>`),
		g.Group(children),
	)
}