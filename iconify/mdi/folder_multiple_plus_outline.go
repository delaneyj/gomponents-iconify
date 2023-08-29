package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMultiplePlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2H6c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h6l2 2h8M2 6v14h18v2H2c-1.1 0-2-.9-2-2V6h2m4 0v10h16V6H6m8 4h2V8h2v2h2v2h-2v2h-2v-2h-2v-2Z"/>`),
		g.Group(children),
	)
}