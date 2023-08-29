package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 1a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 1 0v-9a.5.5 0 0 0-.5-.5Zm-3 1H4v1H2.5a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5H4v1H2.5A1.5 1.5 0 0 1 1 8.5v-5A1.5 1.5 0 0 1 2.5 2Zm6 7H7v1h1.5A1.5 1.5 0 0 0 10 8.5v-5A1.5 1.5 0 0 0 8.5 2H7v1h1.5a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}