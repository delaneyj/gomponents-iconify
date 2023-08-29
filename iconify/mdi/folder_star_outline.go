package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderStarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.78 12.05l3.03-.26L15 9l1.19 2.79l3.03.26l-2.3 1.99l.69 2.96L15 15.47L12.39 17l.69-2.96l-2.3-1.99M22 8v10c0 1.11-.89 2-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h6l2 2h8a2 2 0 0 1 2 2m-2 0H4v10h16V8Z"/>`),
		g.Group(children),
	)
}