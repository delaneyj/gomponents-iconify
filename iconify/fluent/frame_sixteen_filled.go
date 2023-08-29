package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 1a.5.5 0 0 1 .5.5V3h8V1.5a.5.5 0 0 1 1 0V3h1.5a.5.5 0 0 1 0 1H13v8h1.5a.5.5 0 0 1 0 1H13v1.5a.5.5 0 0 1-1 0V13H4v1.5a.5.5 0 0 1-1 0V13H1.5a.5.5 0 0 1 0-1H3V4H1.5a.5.5 0 0 1 0-1H3V1.5a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}