package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretLeftCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0ZM9.5 4.793L6.793 7.5L9.5 10.207l.707-.707l-2-2l2-2l-.707-.707Zm-3 0L3.793 7.5L6.5 10.207l.707-.707l-2-2l2-2l-.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}