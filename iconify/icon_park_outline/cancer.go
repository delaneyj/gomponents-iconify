package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cancer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="37" cy="17" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 13s6-8 16-8s16 6 16 6"/><circle cx="11" cy="31" r="6" transform="rotate(-180 11 31)"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 35s-6 8-16 8s-16-6-16-6"/></g>`),
		g.Group(children),
	)
}