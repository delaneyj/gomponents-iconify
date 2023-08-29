package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 7v10a1 1 0 0 0 1 1h10a1 1 0 0 0 .707-1.707L13.414 12l4.293-4.293a1 1 0 0 0-1.414-1.414L12 10.586L7.707 6.293A1 1 0 0 0 6 7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}