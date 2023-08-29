package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.594 7l-.313.281l-8 8l-.687.719l.687.719l8 8l.313.281H29V7zm.844 2H27v14H12.437l-7-7zm2.718 2.75l-1.406 1.406L16.594 16l-2.844 2.844l1.406 1.406L18 17.406l2.844 2.844l1.406-1.406L19.406 16l2.844-2.844l-1.406-1.406L18 14.594z"/>`),
		g.Group(children),
	)
}