package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineAirlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.34 18H5.8l8.25-12h5.54l-2.25 12zM13 4L2 20h17l3-16h-9zm1.5 5a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5z"/>`),
		g.Group(children),
	)
}