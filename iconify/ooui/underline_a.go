package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnderlineA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 17h14v2H3zm4.7-6.7L10 3.7l2.3 6.6zm6.6 5.7H17L11.5 2h-3L3 16h2.7L7 12h5.8z"/>`),
		g.Group(children),
	)
}