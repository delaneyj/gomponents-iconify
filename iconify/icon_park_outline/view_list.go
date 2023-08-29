package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 14h12M21 24h12M21 34h12"/><path fill="currentColor" fill-rule="evenodd" d="M15 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}