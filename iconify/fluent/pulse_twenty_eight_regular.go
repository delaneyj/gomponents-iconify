package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.023 3a.75.75 0 0 1 .71.59l3.905 17.79l4.148-12.86a.75.75 0 0 1 1.438.033l1.349 4.947h3.677a.75.75 0 0 1 0 1.5H21a.75.75 0 0 1-.724-.553l-.836-3.067l-4.226 13.1a.75.75 0 0 1-1.447-.07L9.905 6.815L7.72 14.456A.75.75 0 0 1 7 15H2.75a.75.75 0 0 1 0-1.5h3.684L9.28 3.544A.75.75 0 0 1 10.023 3Z"/>`),
		g.Group(children),
	)
}