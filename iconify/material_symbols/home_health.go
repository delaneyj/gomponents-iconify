package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17h3v-2.5H16v-3h-2.5V9h-3v2.5H8v3h2.5V17ZM4 21V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}