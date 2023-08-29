package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.812 9.896a1.15 1.15 0 0 1-.65-.197l-6.23-4.156L2.895 9.74a1.175 1.175 0 0 1-1.629-.328a1.173 1.173 0 0 1 .326-1.629L8.28 3.152a1.175 1.175 0 0 1 1.303-.002l6.881 4.59a1.176 1.176 0 0 1-.652 2.156z"/>`),
		g.Group(children),
	)
}