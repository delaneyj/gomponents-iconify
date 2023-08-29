package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20.59L6.59 17H18a2 2 0 0 0 2-2v-.09l-5.21-5.2l-5 5l-2.5-2.5L3 16.5v4.09M20 6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v9.09l4.29-4.3l2.5 2.5l5-5L20 13.5V6M3 22H2V6a3 3 0 0 1 3-3h13a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H7l-4 4M9 6a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2m0 1a1 1 0 0 0-1 1a1 1 0 0 0 1 1a1 1 0 0 0 1-1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}