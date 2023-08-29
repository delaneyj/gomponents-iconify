package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretRightCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm5.5-2.707L8.207 7.5L5.5 10.207L4.793 9.5l2-2l-2-2l.707-.707Zm3 0L11.207 7.5L8.5 10.207L7.793 9.5l2-2l-2-2l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}