package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d0d2d3" d="M63.998 57.1a6.898 6.898 0 0 1-6.897 6.903h-50.2A6.9 6.9 0 0 1 .005 57.1V6.9A6.9 6.9 0 0 1 6.901 0h50.2a6.897 6.897 0 0 1 6.897 6.9v50.2"/>`),
		g.Group(children),
	)
}