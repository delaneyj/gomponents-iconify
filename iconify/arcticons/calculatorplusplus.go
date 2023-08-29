package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculatorplusplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5A21.5 21.5 0 1 1 39.2 8.8L32 16.05a11.25 11.25 0 1 0 0 15.91h0l7.2 7.24A21.46 21.46 0 0 1 24 45.5Zm5.19-25.87v8m-4-4h7.99m3.86 0h8m-4-4v8"/>`),
		g.Group(children),
	)
}