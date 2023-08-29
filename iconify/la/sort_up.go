package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 8.594l-.719.687l-10 10L3.594 21h24.812l-1.687-1.719l-10-10zm0 2.844L23.563 19H8.438z"/>`),
		g.Group(children),
	)
}