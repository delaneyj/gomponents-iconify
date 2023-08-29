package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSplitVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 44H56a12 12 0 0 0-12 12v144a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V56a12 12 0 0 0-12-12ZM56 52h144a4 4 0 0 1 4 4v68H52V56a4 4 0 0 1 4-4Zm144 152H56a4 4 0 0 1-4-4v-68h152v68a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}