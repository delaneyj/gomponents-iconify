package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Textarea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 4c7 0 8 6 8 8v24c0 2-1 8-8 8M32 4c-6 0-8 6-8 8v24c0 2 1 8 8 8M17 24h14"/>`),
		g.Group(children),
	)
}