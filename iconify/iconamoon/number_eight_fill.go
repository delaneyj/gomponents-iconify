package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberEightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 7a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm4.905 2.75a4 4 0 1 0-5.811 0a6 6 0 1 0 5.811 0ZM12 11a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}