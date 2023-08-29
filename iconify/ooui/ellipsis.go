package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="10" cy="10" r="2" fill="currentColor"/><circle cx="3" cy="10" r="2" fill="currentColor"/><circle cx="17" cy="10" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}