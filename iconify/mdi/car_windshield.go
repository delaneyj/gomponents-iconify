package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarWindshield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.63 15.5l-2.42-9.65a1 1 0 0 0-1.07-.75c-1.34.14-5 .4-7.14.4s-5.8-.26-7.14-.4a1 1 0 0 0-1.07.75L1.37 15.5a1.5 1.5 0 0 0 1.18 1.86A61.5 61.5 0 0 0 12 18a61.5 61.5 0 0 0 9.45-.64a1.5 1.5 0 0 0 1.18-1.86Z"/>`),
		g.Group(children),
	)
}