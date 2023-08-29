package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CosmeticLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M11 10.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 20v-4m0 4h3m-3 0h-3"/><path fill="currentColor" d="M2 11v-.75a.75.75 0 0 0-.75.75H2Zm6 0h.75a.75.75 0 0 0-.75-.75V11Zm-6 .75h6v-1.5H2v1.5ZM7.25 11v6h1.5v-6h-1.5Zm-4.5 6v-6h-1.5v6h1.5ZM5 19.25A2.25 2.25 0 0 1 2.75 17h-1.5A3.75 3.75 0 0 0 5 20.75v-1.5ZM7.25 17A2.25 2.25 0 0 1 5 19.25v1.5A3.75 3.75 0 0 0 8.75 17h-1.5Z"/><path stroke="currentColor" stroke-width="1.5" d="M3 11h4V5.618a1 1 0 0 0-1.447-.894l-2 1A1 1 0 0 0 3 6.618V11Z"/></g>`),
		g.Group(children),
	)
}