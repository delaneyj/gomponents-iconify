package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14.17L15.83 12L18 9.83v4.34M20 19V5l-7 7m-9 7h2V5H4m7 9.17L8.83 12L11 9.83v4.34M13 19V5l-7 7"/>`),
		g.Group(children),
	)
}