package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0L0 256l42.7 42.7l64-64V512h298.6V234.7l64 64L512 256z"/>`),
		g.Group(children),
	)
}