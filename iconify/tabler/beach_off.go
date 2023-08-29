package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeachOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.071 15.102a7.502 7.502 0 0 0-8.124 1.648M10.27 6.269L20.196 12a6 6 0 0 0-10.32-6.123"/><path d="M16.732 10C18.39 7.13 18.957 4.356 18 3.804C17.043 3.252 14.925 5.13 13.268 8M15 9l-.739 1.279m-1.467 2.541L12 14.196M3 19.25A2.4 2.4 0 0 1 4 19a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 2-1a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 1.135-.858M3 3l18 18"/></g>`),
		g.Group(children),
	)
}