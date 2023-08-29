package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousMuslim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.096 12.096a6.5 6.5 0 1 1 .48-8.658a5.147 5.147 0 1 0 0 8.124a6.556 6.556 0 0 1-.48.534ZM10.1 6.5H8L9.7 8L9 10.5L11 9l2 1.5l-.7-2.5L14 6.5h-2.1l-.9-2l-.9 2Z"/>`),
		g.Group(children),
	)
}