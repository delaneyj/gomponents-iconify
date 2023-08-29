package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDoneAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 7l-1.41-1.41l-6.34 6.34l1.41 1.41L18 7zm4.24-1.41L11.66 16.17L7.48 12l-1.41 1.41L11.66 19l12-12l-1.42-1.41zM.41 13.41L6 19l1.41-1.41L1.83 12L.41 13.41z"/>`),
		g.Group(children),
	)
}