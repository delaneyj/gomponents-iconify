package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentDividerVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14v7H5V2Zm2 2v3h10V4H7Zm-5 7h20v2H2v-2Zm3 4h14v7H5v-7Zm2 2v3h10v-3H7Z"/>`),
		g.Group(children),
	)
}