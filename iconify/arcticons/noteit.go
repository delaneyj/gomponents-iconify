package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noteit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M40.553 5.5H7.447A1.953 1.953 0 0 0 5.5 7.447v33.106c0 1.07.876 1.947 1.947 1.947h24.888V32.335H42.5V7.447A1.953 1.953 0 0 0 40.553 5.5ZM42.5 32.335L32.335 42.5M5.5 11.809h37"/><path d="m24 19.258l-1.256-1.256a4.249 4.249 0 0 0-5.988 0h0a4.249 4.249 0 0 0 0 5.988l1.256 1.256m11.976 0l1.256-1.256a4.249 4.249 0 0 0 0-5.988h0a4.249 4.249 0 0 0-5.988 0L24 19.258m-5.988 5.988L24 31.235l5.988-5.989"/></g>`),
		g.Group(children),
	)
}