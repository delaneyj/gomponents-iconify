package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restaurant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22v-9.15q-1.275-.35-2.138-1.4T4 9V2h2v7h1V2h2v7h1V2h2v7q0 1.4-.863 2.45T9 12.85V22H7Zm10 0v-8h-3V7q0-2.075 1.463-3.538T19 2v20h-2Z"/>`),
		g.Group(children),
	)
}