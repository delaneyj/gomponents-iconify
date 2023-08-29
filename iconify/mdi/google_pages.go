package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 3h-6v5l4-1l-1 4h5V5a2 2 0 0 0-2-2m-2 14l-4-1v5h6a2 2 0 0 0 2-2v-6h-5m-8 0H3v6a2 2 0 0 0 2 2h6v-5l-4 1M3 5v6h5L7 7l4 1V3H5c-1.11 0-2 .89-2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}