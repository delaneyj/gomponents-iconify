package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExposurePlusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17v-3H3v-2h3V9h2v3h3v2H8v3H6Zm9.75 2V8.05l-2.3 1.65l-1.15-1.75L16.4 5H18v14h-2.25Z"/>`),
		g.Group(children),
	)
}