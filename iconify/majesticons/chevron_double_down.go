package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8 5a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 5H8zm0 8a1 1 0 0 0-.707 1.707l4 4a1 1 0 0 0 1.414 0l4-4A1 1 0 0 0 16 13H8z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}