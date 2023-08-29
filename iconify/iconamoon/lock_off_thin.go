package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOffThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke-linejoin="round" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-9Z"/><path stroke-linejoin="round" stroke-width="1.5" d="M14.5 15.5h.01v.01h-.01z"/></g>`),
		g.Group(children),
	)
}