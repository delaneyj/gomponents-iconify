package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentDuplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6h8a3 3 0 0 1 3 3v11a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3v-1h1v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v5H6V9a3 3 0 0 1 3-3M5 2h10v1H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h7.25L10 13.75l.66-.75l3.5 3.5l-3.5 3.5l-.66-.75L12.25 17H5a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}