package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VennDiagramOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M2.5 9.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path d="M4.5 5.5a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/></g>`),
		g.Group(children),
	)
}