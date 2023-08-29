package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" rx="3"/><path d="M19 19c2 3 2 7 0 10m17-15L21 24l15 10"/><circle cx="16" cy="16" r="4"/><circle cx="16" cy="32" r="4"/></g>`),
		g.Group(children),
	)
}