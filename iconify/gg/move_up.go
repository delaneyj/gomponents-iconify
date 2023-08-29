package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M17 19.071h-2v-8h2v8Zm-8 0H7v-8h2v8Z"/><path d="M13 19.071h-2v-10H7.965l4-4.071l4 4.071H13v10Z"/></g>`),
		g.Group(children),
	)
}