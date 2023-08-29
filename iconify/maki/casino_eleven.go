package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CasinoEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.136 6.85c.29.395.753.65 1.273.65C8.288 7.5 9 6.772 9 5.875c0-.367-.123-.701-.324-.973l.006-.002L5.5 1L2.318 4.9l.006.002c-.2.272-.324.606-.324.973C2 6.772 2.712 7.5 3.59 7.5c.521 0 .983-.255 1.274-.65l.261-.356V8.5c0 .5-1.75.75-1.75.75a.375.375 0 0 0 0 .75h4.25a.375.375 0 0 0 0-.75S5.875 9 5.875 8.5V6.494l.261.356z" fill="currentColor"/>`),
		g.Group(children),
	)
}