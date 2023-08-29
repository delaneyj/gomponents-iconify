package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wunderlist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M17 17.5L12 15l-5 2.5V5H5v14h14V5h-2v12.5m-5-5.08l2.25 1.35l-.6-2.55l1.99-1.72L13 9.27l-1-2.41l-1 2.41l-2.64.23l1.99 1.72l-.6 2.55L12 12.42M5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}