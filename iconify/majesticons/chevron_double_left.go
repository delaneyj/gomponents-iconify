package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 8a1 1 0 0 0-1.707-.707l-4 4a1 1 0 0 0 0 1.414l4 4A1 1 0 0 0 11 16V8zm8 0a1 1 0 0 0-1.707-.707l-4 4a1 1 0 0 0 0 1.414l4 4A1 1 0 0 0 19 16V8z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}