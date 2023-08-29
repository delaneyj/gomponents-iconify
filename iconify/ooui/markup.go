package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 3.5L0 10l1.5 1.5l5 5L8 15l-5-5l5-5zm7 0L12 5l5 5l-5 5l1.5 1.5L20 10z"/>`),
		g.Group(children),
	)
}