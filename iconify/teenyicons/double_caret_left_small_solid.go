package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretLeftSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.207 5.5l-2 2l2 2l-.707.707L3.793 7.5L6.5 4.793l.707.707Zm3 0l-2 2l2 2l-.707.707L6.793 7.5L9.5 4.793l.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}