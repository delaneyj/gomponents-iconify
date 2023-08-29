package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Classroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="13" r="9"/><path d="M5 44c0-8.438 6.175-16.313 11.4-18c0 0 4.75 5.063 7.6 8.438L31.6 26c4.275.563 11.4 9.563 11.4 18"/><path stroke-linecap="round" d="M2 44h44"/></g>`),
		g.Group(children),
	)
}