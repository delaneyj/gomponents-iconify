package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPopulation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 22H17v-2h9v-6h-9v-2h5V6h-5V2h-2v4h-5v6h5v2H6v6h9v2H2v6h13v2h2v-2h13ZM20 8v2h-3V8Zm-8 2V8h3v2Zm12 6v2h-7v-2ZM8 18v-2h7v2Zm-4 8v-2h11v2Zm24 0H17v-2h11Z"/>`),
		g.Group(children),
	)
}