package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDriveSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3.5A1.5 1.5 0 0 1 2.5 2h11A1.5 1.5 0 0 1 15 3.5v2.996a.5.5 0 0 1-.5.5H8.511V8H9.5a.5.5 0 0 1 .5.5V10h1a.5.5 0 0 1 .5.5v1h3a.5.5 0 0 1 0 1h-3v1a.5.5 0 0 1-.5.5H5a.5.5 0 0 1-.5-.5v-1h-3a.5.5 0 0 1 0-1h3v-1A.5.5 0 0 1 5 10h1V8.5a.5.5 0 0 1 .5-.5h1.011V6.996H1.5a.5.5 0 0 1-.5-.5V3.5ZM12.5 5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/>`),
		g.Group(children),
	)
}