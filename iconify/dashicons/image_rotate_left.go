package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageRotateLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5H5.05c0-1.74.85-2.9 2.95-2.9V0C4.85 0 2.96 2.11 2.96 5H1.18L3.8 8.39zm13-4v14h-5v5H1V10h9V1h10zm-2 2h-6v7h3v3h3V3zm-5 9H3v6h10v-6z"/>`),
		g.Group(children),
	)
}