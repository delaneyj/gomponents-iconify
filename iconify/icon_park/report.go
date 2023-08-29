package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Report(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M36 35H12V21C12 14.3726 17.3726 9 24 9C30.6274 9 36 14.3726 36 21V35Z"/><path stroke-linecap="round" d="M8 42H40"/><path stroke-linecap="round" d="M4 13L7 14"/><path stroke-linecap="round" d="M13 3.9999L14 6.9999"/><path stroke-linecap="round" d="M10.0001 9.99989L7.00009 6.99989"/></g>`),
		g.Group(children),
	)
}