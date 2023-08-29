package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="28" x="5" y="14" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 14l14-8m-15 8L10 6m25 14v6"/><rect width="4" height="4" x="33" y="32" fill="currentColor" rx="2"/></g>`),
		g.Group(children),
	)
}