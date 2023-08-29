package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveUpBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5c0-.943 0-1.414.293-1.707C2.586 3 3.057 3 4 3h16c.943 0 1.414 0 1.707.293C22 3.586 22 4.057 22 5c0 .943 0 1.414-.293 1.707C21.414 7 20.943 7 20 7H4c-.943 0-1.414 0-1.707-.293C2 6.414 2 5.943 2 5Zm18.069 3.5c.141 0 .286 0 .431-.002V13c0 3.771 0 5.657-1.172 6.828C18.183 20.974 16.355 21 12.75 21v-7.046l1.693 1.88a.75.75 0 1 0 1.114-1.002l-3-3.334a.75.75 0 0 0-1.114 0l-3 3.334a.75.75 0 1 0 1.114 1.003l1.693-1.88V21c-3.605 0-5.433-.026-6.578-1.172C3.5 18.657 3.5 16.771 3.5 13V8.498c.145.002.29.002.431.002H20.07Z"/>`),
		g.Group(children),
	)
}