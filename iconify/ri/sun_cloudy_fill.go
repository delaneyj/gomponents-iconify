package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunCloudyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.984 5.06a6.5 6.5 0 0 1 11.286 6.436A5.5 5.5 0 0 1 17.5 21H9a8 8 0 1 1 .984-15.941Zm2.071.544a8.026 8.026 0 0 1 4.403 4.495a5.533 5.533 0 0 1 3.12.307a4.5 4.5 0 0 0-7.522-4.802Z"/>`),
		g.Group(children),
	)
}