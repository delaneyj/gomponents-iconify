package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsThreeCircleVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88Zm12-88a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm0-44a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm0 88a12 12 0 1 1-12-12a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}