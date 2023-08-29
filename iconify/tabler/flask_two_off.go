package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.1 15H15m2.742 2.741A6 6 0 0 1 15.318 21H8.683A6 6 0 0 1 10 10.34v-.326M10 6V3h4v7m.613.598a6 6 0 0 1 2.801 2.817M9 3h6M3 3l18 18"/>`),
		g.Group(children),
	)
}