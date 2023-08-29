package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24s6-9 10-9s8 2 10 2s6-2 10-2s10 9 10 9s-10 10-20 10S4 24 4 24Zm0 0h40"/>`),
		g.Group(children),
	)
}