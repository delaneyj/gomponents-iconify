package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageFirst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 16L24 6l1.4 1.4l-8.6 8.6l8.6 8.6L24 26zM8 4h2v24H8z"/>`),
		g.Group(children),
	)
}