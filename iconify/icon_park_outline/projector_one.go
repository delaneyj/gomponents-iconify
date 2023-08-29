package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 21V10H10v11"/><rect width="40" height="16" x="4" y="21" stroke="currentColor" stroke-width="4" rx="2"/><rect width="4" height="4" x="14" y="27" fill="currentColor" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 29h8"/></g>`),
		g.Group(children),
	)
}