package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeTxt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4"/><path d="M14 3v4a1 1 0 0 0 1 1h4m-2.5 7h3"/><path d="M5 12V5a2 2 0 0 1 2-2h7l5 5v4M4.5 15h3M6 15v6m12-6v6m-8-6l4 6m-4 0l4-6"/></g>`),
		g.Group(children),
	)
}