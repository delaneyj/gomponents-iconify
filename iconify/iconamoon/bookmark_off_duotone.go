package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 18l6 3v-3L6 6v15l6-3Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}