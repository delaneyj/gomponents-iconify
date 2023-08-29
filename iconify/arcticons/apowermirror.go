package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apowermirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.979 27.065l7.687-13.314L43.5 34.249H28.168L16.334 13.751L4.5 34.249h12.686"/>`),
		g.Group(children),
	)
}