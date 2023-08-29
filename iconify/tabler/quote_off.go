package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 11H6a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1m4 4v3c0 2.667-1.333 4.333-4 5m13-7h-4m-1-1V7a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v6c0 .66-.082 1.26-.245 1.798m-1.653 2.29c-.571.4-1.272.704-2.102.912M3 3l18 18"/>`),
		g.Group(children),
	)
}