package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeDeciduousTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.998 4.816a4.251 4.251 0 0 1 8.004 0a4.251 4.251 0 0 1 3.364 5.247a4.25 4.25 0 0 1-1.687 7.915a.751.751 0 0 1-.179.022h-3v2.25A1.75 1.75 0 0 1 12.75 22h-1.5a1.75 1.75 0 0 1-1.75-1.75V18h-3a.75.75 0 0 1-.179-.021a4.25 4.25 0 0 1-1.687-7.916a4.256 4.256 0 0 1 3.364-5.247ZM11 18v2.25c0 .138.112.25.25.25h1.5a.25.25 0 0 0 .25-.25V18h-2Z"/>`),
		g.Group(children),
	)
}