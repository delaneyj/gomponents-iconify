package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeaderboardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19h4v-8H4v8Zm6 0h4V5h-4v14Zm6 0h4v-6h-4v6ZM2 21V9h6V3h8v8h6v10H2Z"/>`),
		g.Group(children),
	)
}