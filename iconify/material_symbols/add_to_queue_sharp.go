package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToQueueSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2v-3h3v-2h-3V7h-2v3H8v2h3v3Zm-3 6v-2H2V3h20v16h-6v2H8Z"/>`),
		g.Group(children),
	)
}