package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeedSlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5h6v2H9V5zm10 4h-4V7h4v2zm2 2h-2V9h2v2zm0 6v-6h2v6h-2zm0 0v2h-2v-2h2zM1 11h2v6H1v-6zm2 6h2v2H3v-2zm11-4h-4v-2H8V9H6V7H4v2h2v2h2v2h2v4h4v-4z"/>`),
		g.Group(children),
	)
}