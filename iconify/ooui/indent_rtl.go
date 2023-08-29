package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 16h18v2H1zm0-9h10v2H1zm0 4h10v2H1zm0-9h18v2H1zm18 4v8l-5-4z"/>`),
		g.Group(children),
	)
}