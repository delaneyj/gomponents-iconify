package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M4 42L18 10L28 34L32 22L44 42H4Z"/><path d="M37 14C39.7614 14 42 11.7614 42 9C42 6.23858 39.7614 4 37 4C34.2386 4 32 6.23858 32 9C32 11.7614 34.2386 14 37 14Z"/></g>`),
		g.Group(children),
	)
}