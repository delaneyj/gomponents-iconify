package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watermelon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M23 34C34.0457 34 43 25.0457 43 14H3C3 25.0457 11.9543 34 23 34Z"/><path stroke="#fff" stroke-linecap="round" d="M23 23V26"/><path stroke="#fff" stroke-linecap="round" d="M16.6358 20.3638L14.5145 22.4851"/><path stroke="#fff" stroke-linecap="round" d="M29.3643 20.3642L31.4856 22.4855"/></g>`),
		g.Group(children),
	)
}