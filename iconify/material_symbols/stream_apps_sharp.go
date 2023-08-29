package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamAppsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17V8h9v7h-7l-2 2Zm-9 6V1h14v6h-2V6H7v12h10v-1h2v6H5Z"/>`),
		g.Group(children),
	)
}