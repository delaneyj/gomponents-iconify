package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeedMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5h-2v8h-1v4h4v-4h-1V5zM9 7H5v2H3v2H1v6h2v2h2v-2H3v-6h2V9h4V7zm12 4h2v6h-2v-6zm-2-2h2v2h-2V9zm0 0h-4V7h4v2zm2 8v2h-2v-2h2z"/>`),
		g.Group(children),
	)
}