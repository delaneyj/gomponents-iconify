package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ITwoP(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.758 33.491a15.425 15.425 0 1 1 4.034 1.469"/><circle cx="17.842" cy="17.566" r="4.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.806" cy="16.85" r="4.49" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.181 20.472a1.928 1.928 0 1 1 1.934.82m-10.951.718a1.928 1.928 0 1 1 1.932-.573m-3.338 12.054l-2.827 8.484l4.034 1.469l2.827-8.484"/>`),
		g.Group(children),
	)
}