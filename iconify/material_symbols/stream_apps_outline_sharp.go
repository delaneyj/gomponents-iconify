package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamAppsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13h5v-3h-5v3Zm-2 4V8h9v7h-7l-2 2Zm-9 6V1h14v6h-2V6H7v12h10v-1h2v6H5Zm2-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Zm9-7v-3v3Z"/>`),
		g.Group(children),
	)
}