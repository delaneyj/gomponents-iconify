package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingFourFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20h2v2H1v-2h2V3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v17ZM8 11v2h3v-2H8Zm0-4v2h3V7H8Zm0 8v2h3v-2H8Zm5 0v2h3v-2h-3Zm0-4v2h3v-2h-3Zm0-4v2h3V7h-3Z"/>`),
		g.Group(children),
	)
}