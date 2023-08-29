package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 11a1 1 0 0 0-.707 1.707l7 7a1 1 0 0 0 1.414 0l7-7A1 1 0 0 0 19 11h-6V5a1 1 0 1 0-2 0v6H5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}