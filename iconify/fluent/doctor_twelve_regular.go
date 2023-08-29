package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h1a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H8v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h1V3Zm3 0H5v1.5a.5.5 0 0 1-.5.5H3v2h1.5a.5.5 0 0 1 .5.5V9h2V7.5a.5.5 0 0 1 .5-.5H9V5H7.5a.5.5 0 0 1-.5-.5V3Z"/>`),
		g.Group(children),
	)
}