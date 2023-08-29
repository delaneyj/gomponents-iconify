package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bulldozer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 17a2 2 0 1 0 4 0a2 2 0 0 0-4 0m10 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0m7-4v4a2 2 0 0 0 2 2h1m-8 0H4m0-4h10"/><path d="M9 11V6h2a3 3 0 0 1 3 3v6"/><path d="M5 15v-3a1 1 0 0 1 1-1h8m5 6h-3"/></g>`),
		g.Group(children),
	)
}