package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImageThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2H7Z"/><path stroke-linecap="round" d="m6 20l6-6l6 6"/><path stroke-width="1.5" d="M9.5 10.5h.01v.01H9.5z"/><path d="M13 3v6h6"/></g>`),
		g.Group(children),
	)
}