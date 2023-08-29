package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CategoryManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="14" x="6" y="28" stroke="currentColor" stroke-width="4" rx="4"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M20 7H10a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h10"/><circle cx="34" cy="14" r="8" stroke="currentColor" stroke-width="4"/><circle cx="34" cy="14" r="3" fill="currentColor"/></g>`),
		g.Group(children),
	)
}