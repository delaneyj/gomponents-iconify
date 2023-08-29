package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M16 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/><path d="M8 21.5a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g>`),
		g.Group(children),
	)
}