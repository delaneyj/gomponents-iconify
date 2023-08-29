package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5.51 19.03C3.08 15.235 2.822 7.06 10.551 6.022C15.855 5.46 23.104 15.121 24 42c.896-26.88 8.145-36.54 13.448-35.977c7.729 1.038 7.47 9.213 5.043 13.006"/>`),
		g.Group(children),
	)
}