package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMinus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMinus1" fill="currentColor"><path id="feMinus2" d="M4 11h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2Z"/></g></g>`),
		g.Group(children),
	)
}