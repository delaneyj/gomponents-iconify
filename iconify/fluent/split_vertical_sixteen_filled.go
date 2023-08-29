package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1.5a.5.5 0 0 0-1 0v13a.5.5 0 0 0 1 0v-13ZM1 5a2 2 0 0 1 2-2h3v10H3a2 2 0 0 1-2-2V5Zm8 8h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H9v10Z"/>`),
		g.Group(children),
	)
}