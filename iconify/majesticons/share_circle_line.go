package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 11l-4-4v2c-2 0-6 2.2-6 7c0-.667 1.2-3 6-3v2l4-4z"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}