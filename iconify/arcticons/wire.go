package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 9.55V28.7a11.63 11.63 0 0 0 11.63 11.63h0A11.63 11.63 0 0 0 27.76 28.7V11.42A3.77 3.77 0 0 0 24 7.66h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 9.55V28.7a11.63 11.63 0 0 1-11.63 11.63h0A11.63 11.63 0 0 1 20.24 28.7V11.42A3.77 3.77 0 0 1 24 7.66h0"/>`),
		g.Group(children),
	)
}