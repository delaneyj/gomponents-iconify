package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Muffin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 5s-1-3-4-3s-4 3-4 3C6 5 4 7 4 9c-2.7 0-2.7 4 0 4h16c2.7 0 2.7-4 0-4c0-2-2-4-4-4M5 15l2 7h2l-1-7H5m5 0l1 7h2l1-7h-4m6 0l-1 7h2l2-7h-3Z"/>`),
		g.Group(children),
	)
}