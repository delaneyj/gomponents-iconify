package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8 10h8m-9 4h1m8 0h1"/><path d="M3 18v-6.59a2 2 0 0 1 .162-.787l2.319-5.41A2 2 0 0 1 7.319 4h9.362a2 2 0 0 1 1.838 1.212l2.32 5.41a2 2 0 0 1 .161.789V18M3 18v2.4a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V18m-4 0h4m14 0v2.4a.6.6 0 0 1-.6.6h-2.8a.6.6 0 0 1-.6-.6V18m4 0h-4M7 18h10"/></g>`),
		g.Group(children),
	)
}