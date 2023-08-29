package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Citizenportal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.178 2.5A21.723 21.723 0 1 1 3.479 33.714"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.389 20.983A16.822 16.822 0 1 1 22.8 6.96c.325 0 .646.009.967.027"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.367 13.983l1.19 3.663h3.852l-3.116 2.264l1.19 3.664l-3.116-2.264l-3.116 2.264l1.19-3.664l-3.116-2.264h3.852l1.19-3.663z"/>`),
		g.Group(children),
	)
}