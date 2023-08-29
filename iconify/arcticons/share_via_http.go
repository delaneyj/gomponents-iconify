package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareViaHttp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.344 16.81a6.13 6.13 0 0 0-2.35.468a10.745 10.745 0 0 0-20.274-3.146a7.706 7.706 0 1 0-2.511 14.99h25.135a6.156 6.156 0 1 0 0-12.312Z"/><circle cx="14" cy="22.956" r="3.925" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="22.956" r="3.925" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="35.956" r="3.925" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.608 26.066l-5.216 6.781m-10-6.781l5.216 6.781"/>`),
		g.Group(children),
	)
}