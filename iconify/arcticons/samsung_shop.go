package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.815 15.64l2.056 20.267a8.445 8.445 0 0 0 8.402 7.593h15.454a8.445 8.445 0 0 0 8.402-7.593l2.056-20.268Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.633 15.64v-2.816a8.366 8.366 0 0 1 16.732 0v2.815M20.398 34.394c.8.671 1.663.98 3.602.98h.983a2.899 2.899 0 0 0 2.895-2.902h0a2.899 2.899 0 0 0-2.895-2.902h-1.966a2.898 2.898 0 0 1-2.895-2.902h0a2.899 2.899 0 0 1 2.895-2.901H24c1.94 0 2.803.307 3.602.979"/>`),
		g.Group(children),
	)
}