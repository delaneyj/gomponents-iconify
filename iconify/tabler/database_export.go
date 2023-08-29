package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6c0 1.657 3.582 3 8 3s8-1.343 8-3s-3.582-3-8-3s-8 1.343-8 3"/><path d="M4 6v6c0 1.657 3.582 3 8 3c1.118 0 2.183-.086 3.15-.241M20 12V6"/><path d="M4 12v6c0 1.657 3.582 3 8 3c.157 0 .312-.002.466-.005M16 19h6m-3-3l3 3l-3 3"/></g>`),
		g.Group(children),
	)
}