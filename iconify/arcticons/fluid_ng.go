package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidNg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.162 26.96H17.108a8.27 8.27 0 0 0-8.27 8.27h0a8.27 8.27 0 0 0 8.27 8.27h22.054m-19.324-39h8.324a11 11 0 0 1 11 11v5.54h0H8.838h0V15.5a11 11 0 0 1 11-11Z"/>`),
		g.Group(children),
	)
}