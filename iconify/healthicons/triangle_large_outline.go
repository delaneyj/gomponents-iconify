package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLargeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 5a1 1 0 0 1 .894.553l18 36A1 1 0 0 1 42 43H6a1 1 0 0 1-.894-1.447l18-36A1 1 0 0 1 24 5Zm0 3.236L7.618 41h32.764L24 8.236Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}