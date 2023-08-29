package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m12.781 5.281l-8 8l-.687.719l.687.719l8 8l1.438-1.438L7.938 15H21c2.754 0 5 2.246 5 5v7h2v-7c0-3.844-3.156-7-7-7H7.937l6.282-6.281z"/>`),
		g.Group(children),
	)
}