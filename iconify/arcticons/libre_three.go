package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibreThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.432 43.326a21.486 21.486 0 1 1 9.894-9.893"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.5 12.75a3.5 3.5 0 0 1 7 0v6.5a3.5 3.5 0 0 1-7 0ZM39.15 38.5a2 2 0 0 0 2-2h0a2 2 0 0 0-2-1.999m0 7.999a2 2 0 0 0 2-2h0a2 2 0 0 0-2-2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 41.825a3.407 3.407 0 0 0 2.489.675h.812m-3.301-7.332a3.407 3.407 0 0 1 2.49-.668l.811.002m-2.037 3.999h2.037"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}