package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 30v3a7 7 0 0 1-7 7v0a7 7 0 0 1-7-7V19m24 11v3a7 7 0 0 0 7 7v0a7 7 0 0 0 7-7V19"/><rect width="38" height="22" x="5" y="8" stroke="currentColor" stroke-width="4" rx="11"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 19h-8m4-4v8"/><rect width="4" height="4" x="32" y="15" fill="currentColor" rx="2"/><rect width="4" height="4" x="28" y="20" fill="currentColor" rx="2"/></g>`),
		g.Group(children),
	)
}