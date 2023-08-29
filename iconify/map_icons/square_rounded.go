package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M49 41c0 4.4-3.6 8-8 8H9c-4.4 0-8-3.6-8-8V9c0-4.4 3.6-8 8-8h32c4.4 0 8 3.6 8 8v32z"/>`),
		g.Group(children),
	)
}