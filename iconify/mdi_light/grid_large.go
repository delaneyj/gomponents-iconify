package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4v8h8V6a2 2 0 0 0-2-2h-6m8 9h-8v8h6a2 2 0 0 0 2-2v-6m-9 8v-8H3v6a2 2 0 0 0 2 2h6m-8-9h8V4H5a2 2 0 0 0-2 2v6m2-9h13a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}