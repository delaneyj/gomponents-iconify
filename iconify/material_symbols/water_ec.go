package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterEc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18h1l4-6h-3V8h-1l-4 6h3v4Zm1 4q-3.425 0-5.713-2.35T4 13.8q0-2.5 1.988-5.438T12 2q4.025 3.425 6.013 6.363T20 13.8q0 3.5-2.288 5.85T12 22Z"/>`),
		g.Group(children),
	)
}