package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 1h16v2H0zm6 3h10v2H6zm0 3h10v2H6zm0 3h10v2H6zm-6 3h16v2H0zm0-2V5l4 3z"/>`),
		g.Group(children),
	)
}