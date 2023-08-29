package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BungalowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-6.2l-1.1 1.75l-1.7-1.05L12 3l7.8 12.5l-1.7 1.05L17 14.8V21H7Zm2-2h2v-3h2v3h2v-7.4l-3-4.8l-3 4.8V19Zm2-5v-2h2v2h-2Zm-2 5h6h-6Z"/>`),
		g.Group(children),
	)
}