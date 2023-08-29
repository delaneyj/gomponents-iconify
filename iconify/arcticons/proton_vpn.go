package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.482 34.811l-1.357 2.04a1.323 1.323 0 0 1-2.26-.094l1.398 2.534a3.977 3.977 0 0 0 .929 1.133h0a3.974 3.974 0 0 0 5.864-.873l15.289-23.3a3.969 3.969 0 0 0-2.866-6.12L9.429 6.685a3.969 3.969 0 0 0-3.927 5.861l.119.215l25.578 2.959a2.205 2.205 0 0 1 1.589 3.401ZM5.621 12.761l13.244 23.996"/>`),
		g.Group(children),
	)
}