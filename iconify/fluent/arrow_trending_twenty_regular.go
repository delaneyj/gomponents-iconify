package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.962 5.309A.5.5 0 0 0 17.5 5h-6a.5.5 0 0 0 0 1h4.793L10.5 11.793L8.354 9.646a.5.5 0 0 0-.708 0l-5.5 5.5a.5.5 0 0 0 .708.708L8 10.707l2.146 2.147a.5.5 0 0 0 .708 0L17 6.707V11.5a.5.5 0 0 0 1 0v-6a.5.5 0 0 0-.038-.191Z"/>`),
		g.Group(children),
	)
}