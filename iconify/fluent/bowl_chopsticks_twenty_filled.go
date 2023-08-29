package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlChopsticksTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.238 2.074a.5.5 0 0 1 .688.164L9.087 9h1.826L7.074 2.762a.5.5 0 0 1 .852-.524L12.087 9H17.5a.5.5 0 0 1 .5.5v.5a8 8 0 0 1-.252 2H2.252A8.015 8.015 0 0 1 2 10v-.5a.5.5 0 0 1 .5-.5h5.413L4.074 2.762a.5.5 0 0 1 .164-.688ZM10 18a8.003 8.003 0 0 1-7.418-5h14.837A8.003 8.003 0 0 1 10 18Z"/>`),
		g.Group(children),
	)
}