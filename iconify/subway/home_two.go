package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.7V234.7l64 64L512 256L256 0zm106.7 469.3H149.3V192L256 85.3L362.7 192v277.3z"/>`),
		g.Group(children),
	)
}