package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M44 26L24 36L4 26L24 16L44 26Z"/><path d="M24 7L44 17L24 27L4 17L24 7Z"/><path fill="#2F88FF" stroke-linecap="round" d="M44 26V34L24 44L4 34V26L24 36L44 26Z"/><path stroke-linecap="round" d="M44 14V17V26"/><path stroke-linecap="round" d="M4 26V17V14"/><path stroke-linecap="round" d="M24 36V24"/><path stroke-linecap="round" d="M24 16V4"/></g>`),
		g.Group(children),
	)
}