package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirtyFourHd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 12.549l9.514 25.119L24 10.332l9.986 27.336l9.514-25.12L24 32.399L4.5 12.549Z"/>`),
		g.Group(children),
	)
}