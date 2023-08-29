package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTableBoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.11 0-2 .89-2 2v14c0 1.11.89 2 2 2h14c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2m0 16H5V5h14v14M9 18H6v-2h3v2m4 0h-3v-2h3v2m-4-3H6v-2h3v2m4 0h-3v-2h3v2m-4-3H6v-2h3v2m4 0h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}