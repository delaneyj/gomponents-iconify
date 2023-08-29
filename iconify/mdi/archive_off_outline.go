package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.2 5l-2-2H21v6h-8.8l-2-2H19V5H8.2M20 16.8V10h-2v4.8l2 2m0 2.55v-.01l-2-2v.01L9.66 9l-2-2l-1.53-1.53l-3.74-3.74L1.11 3L3 4.89V9h4.11l10 10H6v-9H4v11h15.11l1.73 1.73l1.27-1.27L20 19.35Z"/>`),
		g.Group(children),
	)
}