package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiddleFinger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M14.972 26.974c-3.32.592-4.981 2.6-4.981 6.026c0 5.138 5.153 11 9.8 11h11.89c4.547 0 7.31-3.85 7.31-6.94V24.01a3 3 0 0 0-3-3h-.01a2.99 2.99 0 0 0-2.99 2.99"/><path d="M14.972 30.04v-9.027a3.003 3.003 0 0 1 3.002-3.003h.003a3.009 3.009 0 0 1 3.006 3.01v4.003"/><path stroke-linejoin="round" d="M20.983 24.008V7.015a3.016 3.016 0 0 1 6.031 0v16.993"/><path stroke-linejoin="round" d="M26.99 23.716v-2.713a3 3 0 0 1 6 0v3"/></g>`),
		g.Group(children),
	)
}