package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BustInSilhouette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M58 61s0-3-1-7c-1.21-4.846-4-8-10-8H25c-6 0-8.79 3.154-10 8c-1 4-1 7-1 7h44zM26 26c0 3.725.54 7.809 2 10c1.861 2.793 5.018 4 8 4c3.096 0 6.137-1.207 8-4c1.46-2.191 2-6.275 2-10c0-2.793-1-12-10-12s-10 7.344-10 12z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M58 60s0-2-1-6c-1.21-4.846-4-8-10-8H25c-6 0-8.79 3.154-10 8c-1 4-1 6-1 6"/><path d="M26 26c0 3.725.54 7.809 2 10c1.861 2.793 5.018 4 8 4c3.096 0 6.137-1.207 8-4c1.46-2.191 2-6.275 2-10c0-2.793-1-12-10-12s-10 7.344-10 12z"/></g>`),
		g.Group(children),
	)
}