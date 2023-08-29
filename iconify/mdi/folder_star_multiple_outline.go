package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderStarMultipleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6v14h18v2H2c-1.105 0-2-.89-2-2V6h2m10.78 4.05l3.03-.26L17 7l1.19 2.79l3.03.26l-2.3 1.99l.69 2.96L17 13.47L14.39 15l.69-2.96l-2.3-1.99M24 6v10c0 1.11-.89 2-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h6l2 2h8a2 2 0 0 1 2 2m-2 0H6v10h16V6Z"/>`),
		g.Group(children),
	)
}