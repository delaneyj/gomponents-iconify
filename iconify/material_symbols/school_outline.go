package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-7-3.8v-6L1 9l11-6l11 6v8h-2v-6.9l-2 1.1v6L12 21Zm0-8.3L18.85 9L12 5.3L5.15 9L12 12.7Zm0 6.025l5-2.7V12.25L12 15l-5-2.75v3.775l5 2.7Zm0-6.025Zm0 2.25Zm0 0Z"/>`),
		g.Group(children),
	)
}