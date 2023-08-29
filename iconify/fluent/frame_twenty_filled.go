package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrameTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2a.5.5 0 0 1 .5.5V4h10V2.5a.5.5 0 0 1 1 0V4h1.5a.5.5 0 0 1 0 1H16v10h1.5a.5.5 0 0 1 0 1H16v1.5a.5.5 0 0 1-1 0V16H5v1.5a.5.5 0 0 1-1 0V16H2.5a.5.5 0 0 1 0-1H4V5H2.5a.5.5 0 0 1 0-1H4V2.5a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}