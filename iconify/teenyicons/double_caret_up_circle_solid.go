package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretUpCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 15a7.5 7.5 0 1 0 0-15a7.5 7.5 0 0 0 0 15Zm2.707-5.5L7.5 6.793L4.793 9.5l.707.707l2-2l2 2l.707-.707Zm0-3L7.5 3.793L4.793 6.5l.707.707l2-2l2 2l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}