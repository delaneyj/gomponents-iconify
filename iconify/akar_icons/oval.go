package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oval(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<ellipse cx="12" cy="12" fill="none" stroke="currentColor" stroke-width="2" rx="8" ry="10"/>`),
		g.Group(children),
	)
}