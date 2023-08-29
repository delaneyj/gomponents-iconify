package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 14.5v-6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M17 9a1 1 0 1 0-2 0a1 1 0 0 0 2 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 11.5a3 3 0 1 0-6 0a3 3 0 0 0 6 0zm-4-7h2a1 1 0 0 1 1 1v1h-4v-1a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}