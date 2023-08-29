package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M28.5 11a.5.5 0 0 0 0-1H21V2.5a.5.5 0 0 0-1 0V10h-9V2.5a.5.5 0 0 0-1 0V10H2.5a.5.5 0 0 0 0 1H10v9H2.5a.5.5 0 0 0 0 1H10v8.5a.5.5 0 0 0 1 0V21h9v8.5a.5.5 0 0 0 1 0V21h7.5a.5.5 0 0 0 0-1H21v-9h7.5zM20 20h-9v-9h9v9z"/><path d="M31 30.5v-29c0-.827-.673-1.5-1.5-1.5h-28C.673 0 0 .673 0 1.5v29c0 .827.673 1.5 1.5 1.5h28c.827 0 1.5-.673 1.5-1.5zm-30 0v-29a.5.5 0 0 1 .5-.5h28a.5.5 0 0 1 .5.5v29a.5.5 0 0 1-.5.5h-28a.5.5 0 0 1-.5-.5z"/></g>`),
		g.Group(children),
	)
}