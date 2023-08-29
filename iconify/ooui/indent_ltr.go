package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 16h18v2H1zm8-9h10v2H9zm0 4h10v2H9zM1 2h18v2H1zm5 8l-5 4V6z"/>`),
		g.Group(children),
	)
}