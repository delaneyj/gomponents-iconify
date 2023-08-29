package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageAltAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16c0 1.103.897 2 2 2h3.586L12 21.414L15.414 18H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v12zM5 4h14v12h-4.414L12 18.586L9.414 16H5V4z"/><path fill="currentColor" d="M11 14h2v-3h3V9h-3V6h-2v3H8v2h3z"/>`),
		g.Group(children),
	)
}