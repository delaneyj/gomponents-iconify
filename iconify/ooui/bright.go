package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bright(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.07 7.07V2.93h-4.14L10 0L7.07 2.93H2.93v4.14L0 10l2.93 2.93v4.14h4.14L10 20l2.93-2.93h4.14v-4.14L20 10zM10 16a6 6 0 1 1 6-6a6 6 0 0 1-6 6z"/><circle cx="10" cy="10" r="4.5" fill="currentColor"/>`),
		g.Group(children),
	)
}