package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypographyOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h3m7 0h6M6.9 15h6.9m-.8-2l3 7M5 20L9.09 9.094m1.091-2.911L11 4h2l3.904 8.924M3 3l18 18"/>`),
		g.Group(children),
	)
}