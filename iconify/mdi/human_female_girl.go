package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumanFemaleGirl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 2a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2M6 22v-6H3l2.6-7.6c.3-.8 1-1.4 1.9-1.4c.9 0 1.7.6 1.9 1.4L12 16H9v6H6m8.5-10a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2m.5 3h3l1.5 4H18v3h-3v-3h-1.5l1.5-4Z"/>`),
		g.Group(children),
	)
}