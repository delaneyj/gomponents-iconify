package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtradStockTrading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.859 25.906l-13.247 5.715l8.069-14.345L35.129 39.69h5.38l-7.226-11.893l4.536-2.452l1.345 1.793l3.586-7.173l-8.069-.896l.745 2.236l-4.331 2.247L21.681 8.31L3.75 39.69l25.408-9.952"/>`),
		g.Group(children),
	)
}