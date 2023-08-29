package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m38 38l-3-3M10 10l3 3"/><path d="M21.143 28L18 17l-3.143 11h6.286Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m14 31l.857-3M22 31l-.857-3m0 0L18 17l-3.143 11m6.286 0h-6.286M35 24c0 5-3.582 7-8 7V17c4.418 0 8 2 8 7Z"/></g>`),
		g.Group(children),
	)
}