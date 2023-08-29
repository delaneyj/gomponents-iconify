package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 11a1 1 0 0 1 .894.553l12 24A1 1 0 0 1 36 37H12a1 1 0 0 1-.894-1.447l12-24A1 1 0 0 1 24 11Zm0 3.236L13.618 35h20.764L24 14.236Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}