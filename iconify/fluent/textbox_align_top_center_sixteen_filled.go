package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxAlignTopCenterSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2A2.5 2.5 0 0 0 2 4.5v7A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 11.5 2h-7Zm1 3h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1 0-1Zm1 3h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}