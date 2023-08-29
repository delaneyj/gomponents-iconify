package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLandscapeTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 10h-4.5A1.5 1.5 0 0 1 12 8.5V4H3.5A1.5 1.5 0 0 0 2 5.5v9A1.5 1.5 0 0 0 3.5 16h13a1.5 1.5 0 0 0 1.5-1.5V10Zm-.25-1H13.5a.5.5 0 0 1-.5-.5V4.25L17.75 9Z"/>`),
		g.Group(children),
	)
}