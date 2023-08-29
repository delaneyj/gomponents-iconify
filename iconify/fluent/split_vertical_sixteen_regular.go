package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM6 13v-1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3V3H3a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h3Zm6 0H9v-1h3a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H9V3h3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}