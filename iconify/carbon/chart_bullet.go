package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 22H16v-1h-2v1H2v6h12v1h2v-1h14zM4 26v-2h10v2zm24 0H16v-2h12zm2-13h-6v-1h-2v1H2v6h20v1h2v-1h6zM4 17v-2h18v2zm24 0h-4v-2h4zm2-13H10V3H8v1H2v6h6v1h2v-1h20zM4 8V6h4v2zm24 0H10V6h18z"/>`),
		g.Group(children),
	)
}