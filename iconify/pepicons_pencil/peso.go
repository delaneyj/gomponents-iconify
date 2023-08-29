package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10.5 2.5h-5v-1h5A4.5 4.5 0 0 1 15 6v1a4.5 4.5 0 0 1-4.5 4.5h-5v-1h5A3.5 3.5 0 0 0 14 7V6a3.5 3.5 0 0 0-3.5-3.5Z"/><path d="M5.5 1.5A.5.5 0 0 1 6 2v15.5a.5.5 0 0 1-1 0V2a.5.5 0 0 1 .5-.5Z"/><path d="M2 4.936a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 0 1h-14a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h14a.5.5 0 0 1 0 1h-14a.5.5 0 0 1-.5-.5Z"/></g>`),
		g.Group(children),
	)
}