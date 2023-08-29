package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleMediumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 8a1 1 0 0 1 .894.553l15 30A1 1 0 0 1 39 40H9a1 1 0 0 1-.894-1.447l15-30A1 1 0 0 1 24 8Zm0 3.236L10.618 38h26.764L24 11.236Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}