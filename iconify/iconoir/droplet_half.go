package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m4.5 16.5l14-6.5m1.5 4c0-4.418-8-12-8-12S4 9.582 4 14a8 8 0 1 0 16 0Z"/>`),
		g.Group(children),
	)
}