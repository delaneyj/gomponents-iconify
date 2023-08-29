package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 18V9C6 7.34315 7.34315 6 9 6H39C40.6569 6 42 7.34315 42 9V18"/><path d="M21 26L33 13"/><path d="M26 21L30 25"/><path d="M31 16L35 20"/><path d="M6 30V39C6 40.6569 7.34315 42 9 42H39C40.6569 42 42 40.6569 42 39V30"/><circle cx="18" cy="30" r="5" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}