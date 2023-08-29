package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveCheckBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 5c0-.943 0-1.414.293-1.707C2.586 3 3.057 3 4 3h16c.943 0 1.414 0 1.707.293C22 3.586 22 4.057 22 5c0 .943 0 1.414-.293 1.707C21.414 7 20.943 7 20 7H4c-.943 0-1.414 0-1.707-.293C2 6.414 2 5.943 2 5Z"/><path fill-rule="evenodd" d="M20.069 8.5c.141 0 .286 0 .431-.002V13c0 3.771 0 5.657-1.172 6.828C18.157 21 16.271 21 12.5 21h-1c-3.771 0-5.657 0-6.828-1.172C3.5 18.657 3.5 16.771 3.5 13V8.498c.145.002.29.002.431.002H20.07Zm-5.07 3.69a.75.75 0 0 1 .06 1.06l-3.571 4a.75.75 0 0 1-1.119 0l-1.428-1.6a.75.75 0 1 1 1.118-1l.87.974l3.012-3.373a.75.75 0 0 1 1.058-.06Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}