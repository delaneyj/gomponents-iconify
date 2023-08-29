package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.707 6.293a1 1 0 0 1 0 1.414L9.414 16H15a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1V9a1 1 0 1 1 2 0v5.586l8.293-8.293a1 1 0 0 1 1.414 0z"/>`),
		g.Group(children),
	)
}