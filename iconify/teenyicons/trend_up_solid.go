package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3h7v7h-1V4.707l-6 6l-3-3l-4.146 4.147l-.708-.708L5 6.293l3 3L13.293 4H8V3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}