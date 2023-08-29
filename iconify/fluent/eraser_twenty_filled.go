package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.197 2.44a1.5 1.5 0 0 1 2.121 0l4.243 4.242a1.5 1.5 0 0 1 0 2.121L9.364 17H14.5a.5.5 0 1 1 0 1H7.82a1.496 1.496 0 0 1-1.14-.437L2.437 13.32a1.5 1.5 0 0 1 0-2.121l8.76-8.76ZM9.781 15.168l-4.95-4.95l-1.687 1.687a.5.5 0 0 0 0 .707l4.243 4.243a.5.5 0 0 0 .707 0l1.687-1.687Z"/>`),
		g.Group(children),
	)
}