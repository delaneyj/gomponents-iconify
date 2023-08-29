package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.46 11.88l1.42-1.42L11 12.59l2.12-2.13l1.42 1.42L12.41 14l2.13 2.12l-1.42 1.42L11 15.41l-2.12 2.13l-1.42-1.42L9.59 14l-2.13-2.12M3 3h18a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2m0 2v14h18V9h-8V5H3Z"/>`),
		g.Group(children),
	)
}