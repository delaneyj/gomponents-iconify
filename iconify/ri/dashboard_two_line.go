package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DashboardTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16Zm0 1c1.018 0 1.985.217 2.858.608L13.295 7.17a5 5 0 0 0-4.83 8.366L7.05 16.95l-.156-.161A7 7 0 0 1 12 5Zm6.392 4.143c.39.872.608 1.84.608 2.857a6.978 6.978 0 0 1-2.05 4.95l-1.414-1.414a5.008 5.008 0 0 0 1.295-4.83l1.561-1.563Zm-2.15-2.8l1.415 1.414l-3.725 3.726A2.003 2.003 0 0 1 12 14a2 2 0 1 1 .517-3.932l3.726-3.725Z"/>`),
		g.Group(children),
	)
}