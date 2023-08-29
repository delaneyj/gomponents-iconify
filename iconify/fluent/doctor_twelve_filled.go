package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoctorTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2a1 1 0 0 0-1 1v1H3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1v1a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V8h1a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H8V3a1 1 0 0 0-1-1H5Z"/>`),
		g.Group(children),
	)
}