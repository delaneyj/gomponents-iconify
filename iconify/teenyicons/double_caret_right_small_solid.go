package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretRightSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 4.793L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2l.707-.707Zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}