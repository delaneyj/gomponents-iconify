package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSplitHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44H56a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM52 200V56a4 4 0 0 1 4-4h68v152H56a4 4 0 0 1-4-4Zm152 0a4 4 0 0 1-4 4h-68V52h68a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}