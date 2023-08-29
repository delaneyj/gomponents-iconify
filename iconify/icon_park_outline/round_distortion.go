package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundDistortion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20"/><path stroke-linecap="round" d="M24 44c-5.523 0-10-4.477-10-10s4.477-10 10-10s10-4.477 10-10S29.523 4 24 4"/><path stroke-linecap="round" d="M44 24c0 5.523-4.477 10-10 10s-10-4.477-10-10s-4.477-10-10-10S4 18.477 4 24"/></g>`),
		g.Group(children),
	)
}