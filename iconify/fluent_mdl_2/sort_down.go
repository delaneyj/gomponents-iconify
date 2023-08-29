package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1364 1459l91 90l-493 493l-494-493l91-90l338 338L896 3h128l1 1795l339-339z"/>`),
		g.Group(children),
	)
}