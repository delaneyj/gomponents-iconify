package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEurThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M187 195a4 4 0 0 1-.31 5.65A76 76 0 0 1 60.11 148H40a4 4 0 0 1 0-8h20v-24H40a4 4 0 0 1 0-8h20.11a76 76 0 0 1 126.56-52.65a4 4 0 1 1-5.34 6A68 68 0 0 0 68.13 108H136a4 4 0 0 1 0 8H68v24h52a4 4 0 0 1 0 8H68.13a68 68 0 0 0 113.2 46.69a4 4 0 0 1 5.67.31Z"/>`),
		g.Group(children),
	)
}