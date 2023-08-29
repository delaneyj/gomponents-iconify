package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 2h18v16l-9 4l-9-4V2Zm14 4H8v5h8v5l-4 1.5L8 16v-2"/>`),
		g.Group(children),
	)
}