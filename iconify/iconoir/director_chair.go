package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectorChair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M19 12L5 21M5 3v9m14-9v9M5 12l14 9M4 12h16"/><path d="M5 4h14M5 7h14"/></g>`),
		g.Group(children),
	)
}