package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2H7Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2H7Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 20l6-6l6 6"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9.5 10.5h.01v.01H9.5z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/></g>`),
		g.Group(children),
	)
}