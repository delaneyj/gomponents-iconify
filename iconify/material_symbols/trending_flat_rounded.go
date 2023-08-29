package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingFlatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.175 13H4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h14.175L16.8 9.6q-.3-.3-.288-.7t.313-.7q.3-.275.713-.287t.687.287l3.075 3.1q.3.3.3.7t-.3.7l-3.1 3.1q-.275.275-.688.275T16.8 15.8q-.3-.3-.3-.713t.3-.712L18.175 13Z"/>`),
		g.Group(children),
	)
}