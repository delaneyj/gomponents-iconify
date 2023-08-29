package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="24" height="32" x="12" y="12" rx="12"/><path d="M12 24c0-6.627 5.373-12 12-12s12 5.373 12 12v1H12v-1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 25V12c0-4 1.5-8 7-8c6 0 8 5 8 9"/></g>`),
		g.Group(children),
	)
}