package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungGalaxyShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.815 15.64l2.056 20.267a8.445 8.445 0 0 0 8.402 7.593h15.454a8.445 8.445 0 0 0 8.402-7.593l2.056-20.268H5.815Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.633 15.64v-2.816a8.366 8.366 0 0 1 16.732 0v2.815"/>`),
		g.Group(children),
	)
}