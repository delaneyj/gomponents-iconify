package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Galaxystore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.766 16.096l1.949 19.208a8.003 8.003 0 0 0 7.962 7.196h14.646a8.003 8.003 0 0 0 7.962-7.196l1.949-19.208Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.071 16.096V13.43a7.929 7.929 0 0 1 15.858 0v2.667"/>`),
		g.Group(children),
	)
}