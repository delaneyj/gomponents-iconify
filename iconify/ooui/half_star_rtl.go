package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HalfStarRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m5.4 12.5l-1.6 7l6.2-3.7l6.2 3.7l-1.6-7L20 7h-7L10 .5L7 7H0zm.8 3.7l1-4.3l-3.7-3.4h4.6L10 4.6v9.3z"/>`),
		g.Group(children),
	)
}