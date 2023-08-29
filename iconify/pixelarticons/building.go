package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18v20H3V2zm12 16v2h4V4H5v16h4v-2h6zM7 6h2v2H7V6zm6 0h-2v2h2V6zm2 0h2v2h-2V6zm-6 4H7v2h2v-2zm2 0h2v2h-2v-2zm6 0h-2v2h2v-2zM7 14h2v2H7v-2zm6 0h-2v2h2v-2zm4 0h-2v2h2v-2z"/>`),
		g.Group(children),
	)
}