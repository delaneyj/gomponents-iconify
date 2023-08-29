package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookMultipleMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M9 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-5v5l-2-1.5L10 7V2H9zM3 6v14a2 2 0 0 0 2 2h12v-2H5V6H3zm16 6v2h-6v-2h6z" fill="currentColor"/>`),
		g.Group(children),
	)
}