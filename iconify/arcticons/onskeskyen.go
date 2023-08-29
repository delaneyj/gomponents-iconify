package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onskeskyen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.353 23.065c-1.278 0-2.463.39-3.446 1.056a4.82 4.82 0 0 0 .62-2.332a4.872 4.872 0 0 0-4.873-4.871c-1.773 0-3.31.957-4.162 2.373a6.826 6.826 0 0 0-13.651.18c0 1.325.397 2.548 1.052 3.594h-2.245a6.147 6.147 0 0 0 0 12.295h26.705a6.147 6.147 0 0 0 0-12.295Z"/>`),
		g.Group(children),
	)
}