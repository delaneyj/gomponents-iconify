package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 16a1 1 0 0 0 1.707.707l4-4a1 1 0 0 0 0-1.414l-4-4A1 1 0 0 0 5 8v8zm8 0a1 1 0 0 0 1.707.707l4-4a1 1 0 0 0 0-1.414l-4-4A1 1 0 0 0 13 8v8z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}