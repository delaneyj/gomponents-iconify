package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorseVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8v8l-3 1l-3.09-5.5c-.26-.46-.99-.23-.91.31L14 21L4 17l1.15-8.06A6.923 6.923 0 0 1 12 3h8l-1.58 2.37C19.36 5.88 20 6.86 20 8Z"/>`),
		g.Group(children),
	)
}