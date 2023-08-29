package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.237 6.237a4.098 4.098 0 0 1 .135 5.654l-7.373 8.11l-7.37-8.11a4.098 4.098 0 1 1 6.23-5.316L12 8l1.14-1.425a4.098 4.098 0 0 1 6.097-.338Z"/>`),
		g.Group(children),
	)
}