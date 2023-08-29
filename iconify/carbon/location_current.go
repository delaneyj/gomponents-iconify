package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationCurrent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 12.983a1 1 0 0 0 .629.945l9.601 3.841l3.841 9.602A1 1 0 0 0 19 28h.016a1 1 0 0 0 .924-.658l8-22a1 1 0 0 0-1.282-1.283l-22 8a1.001 1.001 0 0 0-.658.925Z"/>`),
		g.Group(children),
	)
}