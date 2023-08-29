package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadbarAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><rect width="18" height="4" x="3" y="10" opacity=".3" rx="2"/><rect width="10" height="4" x="7" y="10" rx="2"/></g>`),
		g.Group(children),
	)
}