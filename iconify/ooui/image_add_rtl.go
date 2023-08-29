package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAddRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6v2H8v4H2v6a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zM3.83 17l3.55-4.5l2.52 3l3.55-4.5L18 17zM4 10h2V6h4V4H6V0H4v4H0v2h4z"/>`),
		g.Group(children),
	)
}