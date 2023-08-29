package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10h4v4h-4zm0 0L6.5 6.5M9.96 6A3.5 3.5 0 1 0 6 9.96m8 .04l3.5-3.5m.5 3.46A3.5 3.5 0 1 0 14.04 6M14 14l3.5 3.5m-3.46.5A3.5 3.5 0 1 0 18 14.04M10 14l-3.5 3.5M6 14.04A3.5 3.5 0 1 0 9.96 18"/>`),
		g.Group(children),
	)
}