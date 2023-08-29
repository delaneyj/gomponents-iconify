package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 116h-24a4 4 0 0 0 0 8h19.91A92.13 92.13 0 0 1 132 211.91V91.71a28 28 0 1 0-8 0v120.2A92.13 92.13 0 0 1 36.09 124H56a4 4 0 0 0 0-8H32a4 4 0 0 0-4 4a100 100 0 0 0 200 0a4 4 0 0 0-4-4ZM108 64a20 20 0 1 1 20 20a20 20 0 0 1-20-20Z"/>`),
		g.Group(children),
	)
}