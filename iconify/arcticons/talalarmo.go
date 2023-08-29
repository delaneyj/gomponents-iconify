package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Talalarmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.23 11.22A22.6 22.6 0 0 1 16.92 4.5m14.16 0a22.6 22.6 0 0 1 10.69 6.72M24 8.76h0A17.37 17.37 0 1 1 6.63 26.13A17.35 17.35 0 0 1 24 8.76Zm-.03 17.37V16.11m9.5 18.02l-9.5-8"/>`),
		g.Group(children),
	)
}