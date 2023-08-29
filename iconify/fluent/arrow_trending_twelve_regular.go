package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.962 2.309A.5.5 0 0 0 10.5 2h-4a.5.5 0 0 0 0 1h2.793L6 6.293L4.854 5.146a.5.5 0 0 0-.708 0l-3 3a.5.5 0 1 0 .708.708L4.5 6.207l1.146 1.147a.5.5 0 0 0 .708 0L10 3.707V6.5a.5.5 0 0 0 1 0V2.497a.5.5 0 0 0-.038-.188Z"/>`),
		g.Group(children),
	)
}