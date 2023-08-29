package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddRoad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 23v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2Zm0-10V4h2v9h-2ZM4 20V4h2v16H4Zm7-12V4h2v4h-2Zm0 6v-4h2v4h-2Zm0 6v-4h2v4h-2Z"/>`),
		g.Group(children),
	)
}