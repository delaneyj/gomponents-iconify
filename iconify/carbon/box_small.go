package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V9h2v17h20V9h2v17a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M18 23h-6v-2h6v-4h-4a2.002 2.002 0 0 1-2-2v-4a2.002 2.002 0 0 1 2-2h6v2h-6v4h4a2.002 2.002 0 0 1 2 2v4a2.002 2.002 0 0 1-2 2zM4 4h24v2H4z"/>`),
		g.Group(children),
	)
}