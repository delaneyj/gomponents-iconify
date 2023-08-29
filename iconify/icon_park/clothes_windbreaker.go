package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesWindbreaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M37 25V37M11 37V44H37V37M11 37H4V18L10.125 13L17 18L24 10L31 18L37.875 13L44 18V37H37M11 37V25"/><path fill="#2F88FF" d="M17 18L24 10L31 18L42 10L39.5 4H9L6 10L17 18Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 10H24M29 10H24M24 10L17 18L6 10L9 4H39.5L42 10L31 18L24 10Z"/></g>`),
		g.Group(children),
	)
}