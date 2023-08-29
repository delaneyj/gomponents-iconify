package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3zm0 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM19.5 6H18V2.5a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5V6H4.5A2.502 2.502 0 0 0 2 8.5V15a3.003 3.003 0 0 0 3 3h1v3.5a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5V18h1a3.003 3.003 0 0 0 3-3V8.5A2.502 2.502 0 0 0 19.5 6zM7 3h10v3H7V3zm10 18H7v-6h10v6zm4-6a2.003 2.003 0 0 1-2 2h-1v-2.5a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5V17H5a2.003 2.003 0 0 1-2-2V8.5A1.5 1.5 0 0 1 4.5 7h15A1.5 1.5 0 0 1 21 8.5V15z"/>`),
		g.Group(children),
	)
}