package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaiseToAnswer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="29.426" cy="16.482" r="10.35" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 41.868V31.95a3.881 3.881 0 0 0-3.45-3.45H18.488l-6.9 2.587l6.103-6.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.901 41.868v-6.469l-6.925 1.727c-5.08 1.267-5.498-4.578-2.944-7.606l9.254-10.975m4.054-4.973l-5.175 1.725l4.168 10.764l4.497-1.499"/>`),
		g.Group(children),
	)
}