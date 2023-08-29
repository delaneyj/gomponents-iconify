package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UfoTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M17 7.21c2.989.723 5 2.071 5 3.616C22 13.131 17.523 15 12 15S2 13.13 2 10.826c0-1.545 2.011-2.893 5-3.615"/><path stroke="currentColor" stroke-width="1.5" d="M7 7.729A4.729 4.729 0 0 1 11.729 3h.542A4.729 4.729 0 0 1 17 7.729c0 .177-.054.35-.2.451c-.414.288-1.61.82-4.8.82c-3.19 0-4.386-.532-4.8-.82c-.146-.1-.2-.274-.2-.451Z"/><circle cx="12" cy="12" r="1" fill="currentColor"/><circle cx="7" cy="11" r="1" fill="currentColor"/><circle cx="17" cy="11" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 21v-3m6 2v-3M6 20v-3"/></g>`),
		g.Group(children),
	)
}