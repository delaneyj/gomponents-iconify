package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutdentRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 16h18v2H1zm0-9h10v2H1zm0 4h10v2H1zm0-9h18v2H1zm13 4v8l5-4z"/>`),
		g.Group(children),
	)
}