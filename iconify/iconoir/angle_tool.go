package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleTool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M3 21V3h6v12h12v6H3Z"/><path d="M13 19v2m-4-2v2M3 7h2m-2 4h2m-2 4h2m12 4v2"/></g>`),
		g.Group(children),
	)
}