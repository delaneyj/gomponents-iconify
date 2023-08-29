package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineAirlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4L2 20h17l3-16h-9zm1.5 10a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5z"/>`),
		g.Group(children),
	)
}