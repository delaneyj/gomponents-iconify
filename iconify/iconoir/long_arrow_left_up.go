package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 10.5L8 7l3.5 3.5"/><path d="M8 7v6a4 4 0 0 0 4 4h7"/></g>`),
		g.Group(children),
	)
}