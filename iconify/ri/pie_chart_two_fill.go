package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2.05V13h10.95c-.501 5.053-4.765 9-9.95 9c-5.523 0-10-4.477-10-10c0-5.185 3.947-9.449 9-9.95Zm2-1.507c5.553.478 9.979 4.904 10.457 10.457H13V.543Z"/>`),
		g.Group(children),
	)
}