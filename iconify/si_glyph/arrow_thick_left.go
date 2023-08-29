package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m.133 8.367l5.94 5.346a.648.648 0 0 0 .849-.002l.005-3.728h8.024a.99.99 0 0 0 1-.982V7.035a.99.99 0 0 0-1-.982h-8.02l.005-3.81a.65.65 0 0 0-.848.003L.134 7.603a.503.503 0 0 0-.001.764z"/>`),
		g.Group(children),
	)
}