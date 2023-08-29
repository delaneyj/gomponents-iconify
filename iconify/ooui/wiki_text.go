package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WikiText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3v14h3v-2H3V5h1V3zm4 0v14h4v-2H7V5h2V3zm11 0v2h1v10h-1v2h3V3zm-5 0v2h2v10h-2v2h4V3z"/>`),
		g.Group(children),
	)
}