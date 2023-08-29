package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.39 1.73L1.11 3l1.53 1.53C2.25 4.9 2 5.42 2 6v12a2 2 0 0 0 2 2h14.11l2.73 2.73l1.27-1.27L2.39 1.73M4 18V8h2.11l10 10H4m7.2-10l-4-4H10l2 2h8a2 2 0 0 1 2 2v10c0 .24-.04.47-.12.68L20 16.8V8h-8.8Z"/>`),
		g.Group(children),
	)
}