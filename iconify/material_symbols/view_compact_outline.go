package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCompactOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19.75V4.25h20v15.5H2Zm2-11h2.5v-2.5H4v2.5Zm4.5 0H11v-2.5H8.5v2.5Zm4.5 0h2.5v-2.5H13v2.5Zm4.5 0H20v-2.5h-2.5v2.5Zm0 4.5H20v-2.5h-2.5v2.5Zm-4.5 0h2.5v-2.5H13v2.5Zm-4.5 0H11v-2.5H8.5v2.5Zm-2-2.5H4v2.5h2.5v-2.5Zm11 7H20v-2.5h-2.5v2.5Zm-4.5 0h2.5v-2.5H13v2.5Zm-4.5 0H11v-2.5H8.5v2.5Zm-4.5 0h2.5v-2.5H4v2.5Z"/>`),
		g.Group(children),
	)
}