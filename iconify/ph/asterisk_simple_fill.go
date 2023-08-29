package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm59.15 87.35L141 131.12l33.13 39.76a8 8 0 0 1-12.3 10.24L128 140.5l-33.85 40.62a8 8 0 0 1-12.3-10.24L115 131.12l-46.15-19.77a8 8 0 1 1 6.3-14.7L120 115.87V70a8 8 0 0 1 16 0v45.87l44.85-19.22a8 8 0 1 1 6.3 14.7Z"/>`),
		g.Group(children),
	)
}