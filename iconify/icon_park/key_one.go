package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><circle cx="15" cy="33" r="8" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 16L35.5 22"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 26L37 7"/><path stroke-linecap="round" stroke-linejoin="round" d="M35 11L42 17.5"/></g>`),
		g.Group(children),
	)
}