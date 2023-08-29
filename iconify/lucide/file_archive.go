package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileArchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22V4c0-.5.2-1 .6-1.4C5 2.2 5.5 2 6 2h8.5L20 7.5V20c0 .5-.2 1-.6 1.4c-.4.4-.9.6-1.4.6h-2"/><path d="M14 2v6h6"/><circle cx="10" cy="20" r="2"/><path d="M10 7V6m0 6v-1m0 7v-2"/></g>`),
		g.Group(children),
	)
}