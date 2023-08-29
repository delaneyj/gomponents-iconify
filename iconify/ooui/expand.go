package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m17.5 4.75l-7.5 7.5l-7.5-7.5L1 6.25l9 9l9-9z"/>`),
		g.Group(children),
	)
}