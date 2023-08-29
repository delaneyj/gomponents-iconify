package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleRoundedTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6h6a6 6 0 0 1 6 6v5a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-5a6 6 0 0 1 6-6z"/>`),
		g.Group(children),
	)
}