package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RawFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M4 6V3.6a.6.6 0 0 1 .6-.6h14.8a.6.6 0 0 1 .6.6V6M4 18v2.4a.6.6 0 0 0 .6.6h14.8a.6.6 0 0 0 .6-.6V18"/><path stroke-linejoin="round" d="M16.5 9v6l1.5-3l1.5 3V9m-9 6v-3m0 0v-1.5A1.5 1.5 0 0 1 12 9v0a1.5 1.5 0 0 1 1.5 1.5V12m-3 0h3m0 0v3m-9 0V9h2.4a.6.6 0 0 1 .6.6v.9A1.5 1.5 0 0 1 6 12v0"/><path stroke-linejoin="round" d="M4.5 12H6v0a1.5 1.5 0 0 1 1.5 1.5V15"/></g>`),
		g.Group(children),
	)
}