package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOffLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}