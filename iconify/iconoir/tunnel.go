package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 20L3 14"/><path d="M16 10v1m-4-2v1M8 8v1"/><path stroke-linejoin="round" d="M3 21h18v-9a9 9 0 1 0-18 0v9Z"/></g>`),
		g.Group(children),
	)
}