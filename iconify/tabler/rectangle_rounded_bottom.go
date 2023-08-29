package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleRoundedBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 18h6a6 6 0 0 0 6-6V7a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v5a6 6 0 0 0 6 6z"/>`),
		g.Group(children),
	)
}