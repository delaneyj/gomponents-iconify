package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 8a6 6 0 1 0 0 12A6 6 0 0 0 8 8Zm0 0V3"/><path d="M16 8a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm0 0V3"/></g>`),
		g.Group(children),
	)
}