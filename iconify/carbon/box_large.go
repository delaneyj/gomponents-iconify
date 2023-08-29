package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V9h2v17h20V9h2v17a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M14 21V9h-2v14h8v-2h-6zM4 4h24v2H4z"/>`),
		g.Group(children),
	)
}