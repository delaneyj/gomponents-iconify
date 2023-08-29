package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4v2h1a1 1 0 0 1 1 1h7v2h-7a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1H2v-2h7a1 1 0 0 1 1-1h1v-2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h10m-5 11.5l4.5-4.5H13V6h-2v4H7.5l4.5 4.5Z"/>`),
		g.Group(children),
	)
}