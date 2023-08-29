package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Necktie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m15 36l6-23h6l6 23l-9 8l-9-8Zm6-32h6l3 2l-3 7h-6l-3-7l3-2Z"/>`),
		g.Group(children),
	)
}