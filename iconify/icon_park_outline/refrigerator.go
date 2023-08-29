package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refrigerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="9" y="4" rx="2"/><path stroke-linecap="round" d="M9 22h28M9 20v4m28-4v4m-22 5v4m0-22v4m18 25v4m-20-4v4"/></g>`),
		g.Group(children),
	)
}