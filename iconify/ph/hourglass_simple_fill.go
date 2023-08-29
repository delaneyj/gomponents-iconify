package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M211.31 196.69A16 16 0 0 1 200 224H56a16 16 0 0 1-11.32-27.31a1.59 1.59 0 0 0 .13-.13L116.43 128L44.82 59.44a1.59 1.59 0 0 0-.13-.13A16 16 0 0 1 56 32h144a16 16 0 0 1 11.32 27.31a1.59 1.59 0 0 0-.13.13L139.57 128l71.61 68.56a1.59 1.59 0 0 0 .13.13Z"/>`),
		g.Group(children),
	)
}