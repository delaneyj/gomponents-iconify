package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M5 7h11a8 8 0 0 1 8 8v27a6 6 0 0 0-6-6H5V7Zm38 0H32a8 8 0 0 0-8 8v27a6 6 0 0 1 6-6h13V7Z"/>`),
		g.Group(children),
	)
}